// Package dashscope provides a Go SDK for DashScope Realtime API.
//
// The SDK keeps the public API in the root package and moves protocol,
// authentication, and transport details into internal packages.
//
// Quick start:
//
//	client := dashscope.NewClient(os.Getenv("DASHSCOPE_API_KEY"))
//	session, err := client.Realtime.Connect(context.Background(), &dashscope.RealtimeConfig{
//		Model: dashscope.ModelQwenOmniTurboRealtimeLatest,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer session.Close()
//
//	for event, err := range session.Events() {
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(event.Type)
//	}
package dashscope
