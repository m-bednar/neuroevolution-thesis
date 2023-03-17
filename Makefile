
run:
	go run ./src -env="maps/maze.txt" -pop=200 -maxg=1000 -steps=60 -mutstr=0.2 -tsize=10 -out=output -cmod=20 -nn=3x18

build:
	go build -o neuroevolution.exe ./src 