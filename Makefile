
run:
	go run ./src -env="env.txt" -pop=200 -maxg=2000 -mins=0.99 -steps=35 -mutstr=0.15 -tsize=5 -out=output -cmod=50 -nn=3x16

build:
	go build -o neuroevolution-x86_64-pc-windows-msvc.exe ./src 