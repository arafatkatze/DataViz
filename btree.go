
subgraph cluster_0{
style=filled;color=plum;node [style=filled,color=white, shape="Msquare"];
0[label="1"];1[label="2"];}

	it = tree.Iterator()
	Entries := it.node.Entries
	stringValues := []string{}
	nodeEntrySize := len(Entries)
	dotString := "digraph G{bgcolor=oldlace;subgraph cluster_0 {style=filled;color=plum;node [style=filled,color=white, shape=\"Msquare\"];"
	nodeIndexCount = 0
	for i := 0; i < nodeEntrySize; i++ {
		stringValues = append(stringValues, fmt.Sprintf("%v", Entries[i].Key))
		dotString += strconv.Itoa(nodeIndexCount) + "[label=\"" + stringValues[len(stringValues)-1] + "];"
		nodeIndexCount++
	}
	dotString += "};"
	fmt.Println(dotString)
	for i := 0; it.Next(); i {
		dotString += "subgraph cluster_" + strconv.Itoa(subGraphNumber) + "{style=filled;color=plum;node [style=filled,color=white, shape=\"Msquare\"];"
		Entries = it.node.Entries
		nodeEntrySize = len(Entries)
		for i := 0; i < nodeEntrySize; i++ {
			stringValues = append(stringValues, fmt.Sprintf("%v", Entries[i].Key))
			dotString += strconv.Itoa(nodeIndexCount) + "[label=\"" + stringValues[len(stringValues)-1] + "];"
			nodeIndexCount++
		}
		dotString += "};"
		i += len(Entries)
	}
	fmt.Println(dotString)

	byteString := []byte(dotString) // Converting the string to byte slice to write to a file
	tmpFile, _ := ioutil.TempFile("", "TemporaryDotFile")
	tmpFile.Write(byteString)            // Writing the string to a temporary file
	dotPath, err := exec.LookPath("dot") // Looking for dot command
	if err != nil {
		fmt.Println("Error: Running the Visualizer command. Please install Graphviz")
		return false
	}
	dotCommandResult, err := exec.Command(dotPath, "-Tpng", tmpFile.Name()).Output() // Running the command
	if err != nil {
		fmt.Println("Error: Running the Visualizer command. Please install Graphviz")
		return false
	}
	ioutil.WriteFile("out.png", dotCommandResult, os.FileMode(int(0777)))
