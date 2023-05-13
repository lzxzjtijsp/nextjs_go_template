Docker Imageを作成
```
docker build -t my-go-app . 
```

Docker Imageを実行
```
docker run -p 8080:8080 my-go-app
```
