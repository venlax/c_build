from flask import Flask, request, jsonify
from bisect import bisect_left
from collections import defaultdict
import csv

app = Flask(__name__)


AVAILABLE_TAGS = {}
tag_map = defaultdict(list)



def load_csv(file_path: str):
    global AVAILABLE_TAGS, tag_map

    AVAILABLE_TAGS.clear()
    tag_map.clear()

    with open(file_path, newline="", encoding="utf-8") as csvfile:
        reader = csv.DictReader(csvfile)

        for row in reader:
            tag = row["tag"].strip()
            version = row["version"].strip()

            AVAILABLE_TAGS[version] = tag

            parts = version.split(".")
            if len(parts) == 3:
                key = f"{parts[0]}.{parts[1]}"
                patch = int(parts[2])
                tag_map[key].append((patch, tag))

    for key in tag_map:
        tag_map[key].sort()



def find_best_tag(version: str) -> str | None:
    parts = version.split(".")

    if len(parts) == 2:
        return version

    if len(parts) == 3:
        key = f"{parts[0]}.{parts[1]}"
        patch = int(parts[2])

        if key not in tag_map:
            return None

        patches = tag_map[key]

        idx = bisect_left(patches, (patch, ""))

        if idx < len(patches):
            return patches[idx][1]

        return None

    return None



@app.route("/match", methods=["GET"])
def match_version():
    version = request.args.get("version")

    if not version:
        return jsonify({"error": "version is required"}), 400

    result = find_best_tag(version)

    return jsonify({
        "input": version,
        "match": result
    })



if __name__ == "__main__":
    load_csv("merged.csv")
    app.run(host="0.0.0.0", port=5000)