# go run ./src -env="maps/curve.txt" -pop=300 -maxg=200 -steps=45 -mutstr=0.10 -tsize=30 -out=output -cmod=5 -nn=2x12
# go run ./src -env="maps/maze.txt" -pop=300 -maxg=3000 -steps=60 -mutstr=0.15 -tsize=20 -out=output -cmod=25 -nn=3x24

run:
	go run ./src -env=$(env) -pop=$(pop) -maxg=$(maxg) -steps=$(steps) -mutstr=$(mutstr) -tsize=$(tsize) -out=$(out) -cmod=$(cmod) -nn=$(nn)

build:
	go build -o neuroevolution.exe ./src 
