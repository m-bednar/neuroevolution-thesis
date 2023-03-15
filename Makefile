
run:
	go run ./src -env="env.txt" -pop=200 -maxg=1200 -steps=40 -mutstr=0.15 -tsize=3 -out=output -cmod=5 -nn=3x18

build:
	go build -o neuroevolution-x86_64-pc-windows-msvc.exe ./src 