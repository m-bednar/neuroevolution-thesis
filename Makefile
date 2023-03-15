
run:
	go run ./src -env="env.txt" -pop=200 -maxg=1000 -steps=35 -mutstr=0.15 -tsize=5 -out=output -cmod=10 -nn=3x18

build:
	go build -o neuroevolution-x86_64-pc-windows-msvc.exe ./src 