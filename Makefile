build:
	go mod tidy && go mod vendor
	go build -trimpath -buildmode=plugin -o pluginA.so pluginA.go
	go build -trimpath -buildmode=plugin -o pluginB.so pluginB.go
	go build -trimpath -buildmode=plugin -o pluginBv0_0_2.so pluginBv0_0_2.go
