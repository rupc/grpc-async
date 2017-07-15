# 서버 쪽의 rpc 함수들끼리 공유 변수를 사용해서 통신하는 예제
클라이언트는 3개의 요청을 비동기적을 보냄

1. 클라이언트가 rpc FromClient 를 호출
2. FromClient는 atomic_counter 가 2가 될 때까지 무한 루프에서 대기
3. 클라이언트가 rpc IncreaseCounter 를 호출
4. IncreaseCounter는 atomic_counter 를 1 증가 시킴
5. 클라이언트가 한번 더 rpc IncreaseCounter 를 호출
6. IncreaseCounter는 atomic_counter 를 1 증가 시킴: 2가 됨
7. FromClient 함수는 atomic_counter의 값이 2가 된 것을 보고 루프에서 빠져나옴
8. 그러고나서, 응답 메시지를 클라이언트에게 보냄
9. 클라이언트는 맨처음의 요청을 보낸 고루틴에게 할당된 채널을 통해 응답값을 조사할 수 있음
