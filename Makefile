
run:
	go run ./src -env="env.txt" -pop=50 -maxg=20 -mins=0.99 -steps=20 -mutstr=0.2 -tsize=6 -vout=out.avi

build:
	go build -o neuroevolution-x86_64-pc-windows-msvc.exe ./src 