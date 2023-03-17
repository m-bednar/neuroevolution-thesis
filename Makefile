
run:
	go run ./src -env="env.txt" -pop=200 -maxg=1000 -steps=42 -mutstr=0.2 -tsize=5 -out=output -cmod=10 -nn=3x18

build:
	go build -o neuroevolution.exe ./src 