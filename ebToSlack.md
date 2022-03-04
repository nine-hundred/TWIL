## ebToSlack.go 
### EB에서 발생한 환경 변화를 Slack으로 보내는 Lambda함수입니다.
1. EB에서의 변화를 감지하기 위해 AWS-SNS에서 EB를 "주제"로 등록합니다. 
2. AWS-SNS는 "구독"으로 Lambda를 실행시킵니다. 즉, AWS-SNS는 Lambda의 트리거입니다.
3. 종합하자면, EB -> AWS-SNS -> Lambda -> Slack 인 셈입니다.


```
$ go get -u github.com/slack-go/slack
$ go get -u github.com/aws/aws-lambda-go
$ go get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
```


### 기타
 - 람다에서 AWS-SNS의 정보를 받기 위해선, HandleRequest의 두번째 파라미터를 SNSEvent로 추가하는 것이 좋습니다.

### 참고
- https://docs.aws.amazon.com/ko_kr/lambda/latest/dg/lambda-golang.html
- https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/health-enhanced-status.html
- https://helloinyong.tistory.com/260
- https://jeonghwan-kim.github.io/2018/10/21/aws-sns-to-slack.html
- https://velog.io/@tae2089/Go%EC%97%90%EC%84%9C-Slack-Webhook-%EC%82%AC%EC%9A%A9%ED%95%B4%EB%B3%B4%EA%B8%B0
- https://blog.voidmainvoid.net/221
