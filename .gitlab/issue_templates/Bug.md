**Описание:**  

На stage получилось при регистрации онлайн транзакции возвратить статус 6 (авторизация завершилась ошибкой эквайринга).  
"pass_id": "8174c2e5-d009-4865-976b-48c73c590363"  
Информация о попытке авторизовать транзакцию в таблицу authorizations не записалась, соответственно на доавторизацию транзакция тоже не попала.
В логах кибаны куча ошибок формата MessageFormatError

**Шаги воспроизведения:**  
1. Спуститься вниз на улицу
2. Поздороваться с охранником
3. Позвонить Короткову Д.
4. Завершить 

```json
{
  "pass_id": "4f69e697-b06d-4b07-b5ca-6582de2a67fa"
  "face_id": {
    "id": "1789463",
    "first_name": "Вера",
    "last_name": "Арнольдовна"
  }
}
```  
**Фактический результат:**  
[- по сообщению Ильи - там тестовый стенд вернул-] 1002 ошибку и больше ничего, конечно валидатор взорвался кучей ошибок - поставлю себе в план - если ошибки технического характера от стенда - [-не валидировать 2 -]  

**Ожидаемый результат:**  
@chvalovsl в дожимах нет {+ логов запроса ответа от эквайринга 1 +}    

**Комментарии:**  
[https://kibana.k8s.siroccotechnology.ru/app/kibana#...](https://kibana.k8s.siroccotechnology.ru/app/kibana#/discover?_g=(filters:!(),refreshInterval:(pause:!t,value:0),time:(from:now-60m,to:now))&_a=(columns:!(kubernetes.labels.app,msg,trace-id,kubernetes.namespace_labels.name,passId,traceId),filters:!(('$state':(store:appState),meta:(alias:!n,disabled:!f,index:'92ee2660-d57d-11e9-bbe3-ed404befa7a2',key:kubernetes.labels.app,negate:!f,params:(query:abs-service),type:phrase,value:abs-service),query:(match:(kubernetes.labels.app:(query:abs-service,type:phrase))))),index:'92ee2660-d57d-11e9-bbe3-ed404befa7a2',interval:auto,query:(language:kuery,query:''),sort:!('@timestamp',desc)))  
[https://kibana.k8s.siroccotechnology.ru/app/kibana#...](https://kibana.k8s.siroccotechnology.ru/app/kibana#/discover?_g=(filters:!(),refreshInterval:(pause:!t,value:0),time:(from:now-60m,to:now))&_a=(columns:!(kubernetes.labels.app,msg,trace-id,kubernetes.namespace_labels.name,passId,traceId),filters:!(('$state':(store:appState),meta:(alias:!n,disabled:!f,index:'92ee2660-d57d-11e9-bbe3-ed404befa7a2',key:kubernetes.labels.app,negate:!f,params:(query:abs-service),type:phrase,value:abs-service),query:(match:(kubernetes.labels.app:(query:abs-service,type:phrase))))),index:'92ee2660-d57d-11e9-bbe3-ed404befa7a2',interval:auto,query:(language:kuery,query:''),sort:!('@timestamp',desc)))

**Пруфы:**  
![alt text](/uploads/ae7fb35f22570186a9d569f4e6b12cc5/123.png "Title Text")  

@its.crowd 
@chvalovsl 
@aleksashkametelkin 