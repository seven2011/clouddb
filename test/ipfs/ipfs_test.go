package ipfs

import (
	"testing"
)

func TestIpfsNode(t *testing.T){

	//InitIpfs()

}

/// ------ Setting up the IPFS Repo

//func setupPlugins(externalPluginsPath string) error {
//	// Load any external plugins if available on externalPluginsPath
//	plugins, err := loader.NewPluginLoader(filepath.Join(externalPluginsPath, "plugins"))
//	if err != nil {
//		return fmt.Errorf("error loading plugins: %s", err)
//	}
//	// Load preloaded and external plugins
//	if err := plugins.Initialize(); err != nil {
//		return fmt.Errorf("error initializing plugins: %s", err)
//	}
//	if err := plugins.Inject(); err != nil {
//		return fmt.Errorf("error initializing plugins: %s", err)
//	}
//	return nil
//}
//
//func createTempRepo(ctx context.Context) (string, error, string) {
//	//repoPath, err := ioutil.TempDir("", "example-folder-shell")
//	repoPath := "./example-folder/.example-folder"
//	//if err != nil {
//	//	return "", fmt.Errorf("failed to get temp dir: %s", err)
//	//}
//
//	// Create a config with default options and a 2048 bit key
//	cfg, err := config.Init(ioutil.Discard, 2048)
//	if err != nil {
//		return "", err, ""
//	}
//
//	peerId := cfg.Identity.PeerID
//
//	// Create the repo with the config
//	err = fsrepo.Init(repoPath, cfg)
//	if err != nil {
//		return "", fmt.Errorf("failed to init ephemeral node: %s", err), ""
//	}
//
//	return repoPath, nil, peerId
//}
//
///// ------ Spawning the node
//
//// Creates an IPFS node and returns its coreAPI
//func createNode(ctx context.Context, repoPath string) (icore.CoreAPI, error) {
//	// Open the repo
//	repo, err := fsrepo.Open(repoPath)
//	if err != nil {
//		return nil, err
//	}
//
//	// Construct the node
//
//	nodeOptions := &core.BuildCfg{
//		Online:  true,
//		Routing: libp2p.DHTOption, // This option sets the node to be a full DHT node (both fetching and storing DHT Records)
//		// Routing: libp2p.DHTClientOption, // This option sets the node to be a client DHT node (only fetching records)
//		Repo: repo,
//		ExtraOpts: map[string]bool{
//			"pubsub": true,
//		},
//	}
//
//	node, err := core.NewNode(ctx, nodeOptions)
//	if err != nil {
//		return nil, err
//	}
//
//	// Attach the Core API to the constructed node
//	return coreapi.NewCoreAPI(node)
//}
//
//// Spawns a node on the default repo location, if the repo exists
//func spawnDefault(ctx context.Context) (icore.CoreAPI, error) {
//	defaultPath, err := config.PathRoot()
//	if err != nil {
//		// shouldn't be possible
//		return nil, err
//	}
//
//	if err := setupPlugins(defaultPath); err != nil {
//		return nil, err
//
//	}
//
//	return createNode(ctx, defaultPath)
//}
//
//// Spawns a node to be used just for this run (i.e. creates a tmp repo)
//func spawnEphemeral(ctx context.Context) (icore.CoreAPI, error, string) {
//	if err := setupPlugins(""); err != nil {
//		return nil, err, ""
//	}
//
//	// Create a Temporary Repo
//
//	repoPath, err, pid := createTempRepo(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("failed to create temp repo: %s", err), ""
//	}
//
//	// Spawning an ephemeral IPFS node
//	i, e := createNode(ctx, repoPath)
//	return i, e, pid
//}
//
////
//
//func connectToPeers(ctx context.Context, ipfs icore.CoreAPI, peers []string) error {
//	var wg sync.WaitGroup
//	peerInfos := make(map[peer.ID]*peerstore.PeerInfo, len(peers))
//	for _, addrStr := range peers {
//		addr, err := ma.NewMultiaddr(addrStr)
//		if err != nil {
//			return err
//		}
//		pii, err := peerstore.InfoFromP2pAddr(addr)
//		if err != nil {
//			return err
//		}
//		pi, ok := peerInfos[pii.ID]
//		if !ok {
//			pi = &peerstore.PeerInfo{ID: pii.ID}
//			peerInfos[pi.ID] = pi
//		}
//		pi.Addrs = append(pi.Addrs, pii.Addrs...)
//	}
//
//	wg.Add(len(peerInfos))
//	for _, peerInfo := range peerInfos {
//		go func(peerInfo *peerstore.PeerInfo) {
//			defer wg.Done()
//			err := ipfs.Swarm().Connect(ctx, *peerInfo)
//			if err != nil {
//				log.Printf("failed to connect to %s: %s", peerInfo.ID, err)
//			}
//		}(peerInfo)
//	}
//	wg.Wait()
//	return nil
//}
//
//func getUnixfsFile(path string) (files.File, error) {
//	file, err := os.Open(path)
//	if err != nil {
//		return nil, err
//	}
//	defer file.Close()
//
//	st, err := file.Stat()
//	if err != nil {
//		return nil, err
//	}
//
//	f, err := files.NewReaderPathFile(path, file, st)
//	if err != nil {
//		return nil, err
//	}
//
//	return f, nil
//}
//
//func getUnixfsNode(path string) (files.Node, error) {
//	st, err := os.Stat(path)
//	if err != nil {
//		return nil, err
//	}
//
//	f, err := files.NewSerialFile(path, false, st)
//	if err != nil {
//		return nil, err
//	}
//
//	return f, nil
//}
//
///// -------
//
//func InitIpfs() {
//	/// --- Part I: Getting a IPFS node running
//
//	fmt.Println("-- Getting an IPFS node running -- ")
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	/*
//		// Spawn a node using the default path (~/.example-folder), assuming that a repo exists there already
//		fmt.Println("Spawning node on default repo")
//		example-folder, err := spawnDefault(ctx)
//		if err != nil {
//			fmt.Println("No IPFS repo available on the default path")
//		}
//	*/
//
//	// Spawn a node using a temporary path, creating a temporary repo for the run
//	fmt.Println("Spawning node on a temporary repo")
//
//	ipfs, err, peerId := spawnEphemeral(ctx)
//	ipfs.PubSub()
//	fmt.Println("这是 IPFS 节点信息 ", ipfs)
//	fmt.Println("这是 IPFS peerId信息 ", peerId)
//
//	api := ipfs.Name()
//
//	fmt.Println("打印 core api =", api)
//
//	if err != nil {
//		panic(fmt.Errorf("failed to spawn ephemeral node: %s", err))
//	}
//
//	fmt.Println("IPFS node is running")
//
//	/// --- Part II: Adding a file and a directory to IPFS
//
//	fmt.Println("\n-- Adding and getting back files & directories --")
//
//	inputBasePath := "./example-folder.txt"
//	inputPathFile := inputBasePath
//	inputPathDirectory := inputBasePath
//
//	someFile, err := getUnixfsNode(inputPathFile)
//	if err != nil {
//		panic(fmt.Errorf("Could not get File: %s", err))
//	}
//
//	cidFile, err := ipfs.Unixfs().Add(ctx, someFile)
//	if err != nil {
//		panic(fmt.Errorf("Could not add File: %s", err))
//	}
//
//	fmt.Printf("Added file to IPFS with CID %s\n", cidFile.String())
//
//	someDirectory, err := getUnixfsNode(inputPathDirectory)
//	if err != nil {
//		panic(fmt.Errorf("Could not get File: %s", err))
//	}
//
//	cidDirectory, err := ipfs.Unixfs().Add(ctx, someDirectory)
//	if err != nil {
//		panic(fmt.Errorf("Could not add Directory: %s", err))
//	}
//
//	fmt.Printf("Added directory to IPFS with CID %s\n", cidDirectory.String())
//
//	/// --- Part III: Getting the file and directory you added back
//
//	outputBasePath := "./output/output.txt"
//	outputPathFile := outputBasePath + strings.Split(cidFile.String(), "/")[2]
//	split := strings.Split(cidFile.String(), "/")[2]
//	fmt.Println(" 这是 split =", split)
//
//	fmt.Println(" 这是 outputBasePath =", outputBasePath)
//	fmt.Println(" 这是 outputPathFile =", outputPathFile)
//
//	outputPathDirectory := outputBasePath + strings.Split(cidDirectory.String(), "/")[2]
//
//	rootNodeFile, err := ipfs.Unixfs().Get(ctx, cidFile)
//	if err != nil {
//		panic(fmt.Errorf("Could not get file with CID: %s", err))
//	}
//
//	err = files.WriteTo(rootNodeFile, outputPathFile)
//	if err != nil {
//		panic(fmt.Errorf("Could not write out the fetched CID: %s", err))
//	}
//
//	fmt.Printf("Got file back from IPFS (IPFS path: %s) and wrote it to %s\n", cidFile.String(), outputPathFile)
//
//	rootNodeDirectory, err := ipfs.Unixfs().Get(ctx, cidDirectory)
//	if err != nil {
//		panic(fmt.Errorf("Could not get file with CID: %s", err))
//	}
//
//	err = files.WriteTo(rootNodeDirectory, outputPathDirectory)
//	if err != nil {
//		panic(fmt.Errorf("Could not write out the fetched CID: %s", err))
//	}
//
//	fmt.Printf("Got directory back from IPFS (IPFS path: %s) and wrote it to %s\n", cidDirectory.String(), outputPathDirectory)
//
//	/// --- Part IV: Getting a file from the IPFS Network
//
//	fmt.Println("\n-- Going to connect to a few nodes in the Network as bootstrappers --")
//
//	bootstrapNodes := []string{
//		// IPFS Bootstrapper nodes.
//		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
//		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
//		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
//		"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
//
//		// IPFS Cluster Pinning nodes
//		"/ip4/138.201.67.219/tcp/4001/p2p/QmUd6zHcbkbcs7SMxwLs48qZVX3vpcM8errYS7xEczwRMA",
//		"/ip4/138.201.67.219/udp/4001/quic/p2p/QmUd6zHcbkbcs7SMxwLs48qZVX3vpcM8errYS7xEczwRMA",
//		"/ip4/138.201.67.220/tcp/4001/p2p/QmNSYxZAiJHeLdkBg38roksAR9So7Y5eojks1yjEcUtZ7i",
//		"/ip4/138.201.67.220/udp/4001/quic/p2p/QmNSYxZAiJHeLdkBg38roksAR9So7Y5eojks1yjEcUtZ7i",
//		"/ip4/138.201.68.74/tcp/4001/p2p/QmdnXwLrC8p1ueiq2Qya8joNvk3TVVDAut7PrikmZwubtR",
//		"/ip4/138.201.68.74/udp/4001/quic/p2p/QmdnXwLrC8p1ueiq2Qya8joNvk3TVVDAut7PrikmZwubtR",
//		"/ip4/94.130.135.167/tcp/4001/p2p/QmUEMvxS2e7iDrereVYc5SWPauXPyNwxcy9BXZrC1QTcHE",
//		"/ip4/94.130.135.167/udp/4001/quic/p2p/QmUEMvxS2e7iDrereVYc5SWPauXPyNwxcy9BXZrC1QTcHE",
//
//
//		// You can add more nodes here, for example, another IPFS node you might have running locally, mine was:
//		// "/ip4/127.0.0.1/tcp/4010/p2p/QmZp2fhDLxjYue2RiUvLwT9MWdnbDxam32qYFnGmxZDh5L",
//		// "/ip4/127.0.0.1/udp/4010/quic/p2p/QmZp2fhDLxjYue2RiUvLwT9MWdnbDxam32qYFnGmxZDh5L",
//	}
//	/*
//	"/ip4/138.201.67.219/tcp/4001/p2p/QmUd6zHcbkbcs7SMxwLs48qZVX3vpcM8errYS7xEczwRMA",
//			"/ip4/138.201.67.219/udp/4001/quic/p2p/QmUd6zHcbkbcs7SMxwLs48qZVX3vpcM8errYS7xEczwRMA",
//			"/ip4/138.201.67.220/tcp/4001/p2p/QmNSYxZAiJHeLdkBg38roksAR9So7Y5eojks1yjEcUtZ7i",
//			"/ip4/138.201.67.220/udp/4001/quic/p2p/QmNSYxZAiJHeLdkBg38roksAR9So7Y5eojks1yjEcUtZ7i",
//			"/ip4/138.201.68.74/tcp/4001/p2p/QmdnXwLrC8p1ueiq2Qya8joNvk3TVVDAut7PrikmZwubtR",
//			"/ip4/138.201.68.74/udp/4001/quic/p2p/QmdnXwLrC8p1ueiq2Qya8joNvk3TVVDAut7PrikmZwubtR",
//			"/ip4/94.130.135.167/tcp/4001/p2p/QmUEMvxS2e7iDrereVYc5SWPauXPyNwxcy9BXZrC1QTcHE",
//			"/ip4/94.130.135.167/udp/4001/quic/p2p/QmUEMvxS2e7iDrereVYc5SWPauXPyNwxcy9BXZrC1QTcHE",
//			"/ip4/192.168.120.85/tcp/4001/p2p/12D3KooWGUmM7hiv98hBr5NMyCnHkShk7qe6xnBci94MvBJ9u9dH",
//			"/ip4/182.150.116.150/tcp/4001/p2p/12D3KooWGUmM7hiv98hBr5NMyCnHkShk7qe6xnBci94MvBJ9u9dH",
//	 */
//	go connectToPeers(ctx, ipfs, bootstrapNodes)
//
//	exampleCIDStr := "QmUaoioqU7bxezBQZkUcgcSyokatMY71sxsALxQmRRrHrj"
//
//	fmt.Printf("Fetching a file from the network with CID %s\n", exampleCIDStr)
//	outputPath := outputBasePath + exampleCIDStr
//	testCID := icorepath.New(exampleCIDStr)
//
//	rootNode, err := ipfs.Unixfs().Get(ctx, testCID)
//	if err != nil {
//		panic(fmt.Errorf("Could not get file with CID: %s", err))
//	}
//
//	err = files.WriteTo(rootNode, outputPath)
//	if err != nil {
//		panic(fmt.Errorf("Could not write out the fetched CID: %s", err))
//	}
//
//	fmt.Printf("Wrote the file to %s\n", outputPath)
//
//	fmt.Println("\nAll done! You just finalized your first tutorial on how to use go-example-folder as a library")
//	//fmt.Println("======   开始创建 orbit 数据库 ====")
//	//var dbPath string = "/Users/apple/Desktop/orbit/db"
//	////dbPath, _ := ioutil.TempDir("", "db-keystore")
//	//fmt.Println("dbPath =", dbPath)
//	//
//	//ac := &accesscontroller.CreateAccessControllerOptions{
//	//	Access: map[string][]string{
//	//		"write": {
//	//			"*",
//	//		},
//	//	},
//	//}
//	//orbit, err := orbitdb.NewOrbitDB(ctx, example-folder, &orbitdb.NewOrbitDBOptions{Directory: &dbPath})
//	//fmt.Println("orbit =", orbit)
//	//fmt.Println("创建 db 数据库 ")
//	/*
//
//			//db, err := orbit.Create(ctx, "f1", "keyvalue", nil)
//
//		db, err := orbit.KeyValue(ctx, "f1", &orbitdb.CreateDBOptions{
//			AccessController: ac,
//		})
//		if err != nil {
//			fmt.Println("这是 orbit 的数据 err = ", err)
//
//		}
//		db1, err1 := orbit.Log(ctx, "log database", nil)
//		fmt.Println("这是 db1 的数据 db1 = ", db1)
//
//		if err1 != nil {
//			fmt.Println("这是 orbit 的数据 err = ", err)
//
//		}
//		fmt.Println("这是 orbit 的数据 类型 = ", db.Type())
//		fmt.Println("这是 orbit 的数据 类型 = ", db.Identity())
//
//		fmt.Println("创建 orbit 数据库 类型 ", orbit)
//
//		fmt.Println("创建 orbit 数据库 类型 ", orbit.Identity().Type)
//		fmt.Println("创建 orbit 数据库 id ", orbit.Identity().ID)
//		fmt.Println("创建 orbit 数据库 public", string(orbit.Identity().PublicKey))
//		fmt.Println("创建 orbit 数据库 signatures =", string(orbit.Identity().Signatures.PublicKey))
//		fmt.Println("创建 orbit 数据库 signatures =", string(orbit.Identity().Signatures.ID))
//
//		s, err := db.Put(ctx, "key1", []byte("hello4"))
//		if err != nil {
//			fmt.Println("err = put", err)
//		}
//		fmt.Println("s = ", s)
//
//		value, err := db.Get(ctx, "key1")
//		fmt.Println("value = ", string(value))
//	*/
//	//======== 测试 Docs 数据库
//	/*
//		db2, err := orbit.Docs(ctx, "f2", &orbitdb.CreateDBOptions{
//			AccessController: ac,
//		})
//		fmt.Println(" db2 的数据 类型", db2.Type())
//	*/
//	//db3, err := orbit.KeyValue(ctx, "f3", &orbitdb.CreateDBOptions{
//	//	AccessController: ac,
//	//})
//	////fmt.Println("这是 测试的  db2 =====", db2)
//	//fmt.Println("这是 测试的  db3 =====", db3)
//	//
//	//fmt.Println("-------------------------  这是 测试的  db2 -----------------------")
//	//
//	//l, err := db3.Put(ctx, "t1", []byte("你好"))
//	//if err != nil {
//	//	fmt.Println("这是 测试的  db2 err =====", err)
//	//
//	//}
//	//fmt.Println("这是 测试的  l =====", l)
//	//
//	//l2, _ := db3.Get(ctx, "t1")
//	//fmt.Println("这是 测试的  l2 =====", string(l2))
//	//
//	//l1, err := db3.Put(ctx, "s", []byte("你好a !"))
//	//if err != nil {
//	//	fmt.Println("这是 测试的  db2 err =====", err)
//	//
//	//}
//	//fmt.Println("这是 测试的  l =====", l1)
//	//
//	//l3, _ := db3.Get(ctx, "s")
//	//fmt.Println("这是 测试的  l2 =====", string(l3))
//
//	//example-folder  pub
//	ipfs.PubSub().Publish(ctx, "fly", []byte("于欢 喜欢  抽烟"))
//
//	time.Sleep(1 * time.Hour)
//}
