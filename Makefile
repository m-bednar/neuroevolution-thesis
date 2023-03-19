
run:
	go run ./src -env="maps/maze.txt" -pop=300 -maxg=5000 -steps=60 -mutstr=0.15 -tsize=20 -out=output -cmod=25 -nn=3x24

build:
	go build -o neuroevolution.exe ./src 