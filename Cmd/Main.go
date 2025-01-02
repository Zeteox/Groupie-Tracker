package main

import internal "GroupieTracker/Internal/Server"

func main() {
	internal.CreateAndListenServer(8080)
}
