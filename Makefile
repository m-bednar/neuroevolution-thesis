
maze:
	go run ./src -env="maps/maze.txt" -pop=300 -maxg=5000 -steps=60 -mutstr=0.15 -tsize=20 -out=output -cmod=25 -nn=3x24

curve:
	go run ./src -env="maps/curve.txt" -pop=300 -maxg=1500 -steps=45 -mutstr=0.15 -tsize=30 -out=output -cmod=20 -nn=2x12

build:
	go build -o neuroevolution.exe ./src 