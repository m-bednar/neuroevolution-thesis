
run:
	go run ./src -env="env.txt" -pop=200 -maxg=100 -mins=0.99 -steps=20 -mutstr=0.2 -tsize=6 -out=output -cmod=2

build:
	go build -o neuroevolution-x86_64-pc-windows-msvc.exe ./src 