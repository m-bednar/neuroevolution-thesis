module github.com/m-bednar/neuroevolution-thesis

go 1.19

require github.com/llgcode/draw2d v0.0.0-20210904075650-80aa0a2a901d

require (
	github.com/fzipp/astar v0.2.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/icza/mjpeg v0.0.0-20220812133530-f79265a232f2 // indirect
	github.com/myfantasy/mft v0.2.3 // indirect
	github.com/wcharczuk/go-chart v2.0.1+incompatible // indirect
	golang.org/x/image v0.6.0 // indirect
)

replace (
	github.com/m-bednar/neuroevolution-thesis/src/enviroment => ./src/enviroment
	github.com/m-bednar/neuroevolution-thesis/src/microbe => ./src/microbe
	github.com/m-bednar/neuroevolution-thesis/src/neuralnet => ./src/neuralnet
	github.com/m-bednar/neuroevolution-thesis/src/output => ./src/output
	github.com/m-bednar/neuroevolution-thesis/src/strategies => ./src/strategies
	github.com/m-bednar/neuroevolution-thesis/src/utils => ./src/utils
)
