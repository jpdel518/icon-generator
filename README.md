## Icon Generator

### Docker
#### build
Docker build -t icon_generator .

#### run
Docker run -p 8080:8080 -d icon_generator

### 生成リクエスト
#### Letter Icon Generator
http://localhost:8080/icon?letter=あ&size=125  
- letter: アイコンにしたい文字  
- size: アイコンサイズ  

#### Identicon Generator
http://localhost:8080/identicon?letter=あ&size=125  
- letter: Identiconのseedになる文字
- size: アイコンサイズ
