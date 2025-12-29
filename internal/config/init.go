package config

func Init() {
	PkgMgrName = "apt"
	Image = "ubuntu:22.04"
	// Image = "gcc:13"
	Libs = append(Libs, LibInfo{
		// Name : "tmux",
		// Version: "3.2a-4ubuntu0.2",
		Name : "build-essential",
		Version: "12.9ubuntu3",
	})
}