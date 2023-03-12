
run:
	go run ./src -env="env.txt" -pop=200 -maxg=500 -mins=0.99 -steps=30 -mutstr=0.2 -tsize=6 -out=output -cmod=5 -nn=2x8

build:
	go build -o neuroevolution-x86_64-pc-windows-msvc.exe ./src 