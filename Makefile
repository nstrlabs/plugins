build:
	go build -buildmode=plugin -o pluginA.so pluginA.go
	go build -buildmode=plugin -o pluginB.so pluginB.go
	go build -buildmode=plugin -o pluginBv0_0_2.so pluginBv0_0_2.go
