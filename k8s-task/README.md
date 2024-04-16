# Домашнее задание по docker compose



## Сборка и запуск
Скачайте репозиторий приложения
```
git clone https://github.com/vanc0uv3r/cloud-computing-course.git && cd cloud-computing-course/k8s-task
```
Убедитесь, что на вашей машине установлены **kubectl_** и **_minikube_**        
Для **сборки** используйте команду:  
```
chmod +x setup-manifests.sh && ./setup-manifests.sh
```      

Запросим адрес приложения
```
minikube service go-server  --url
```

Также с помощью команды 
```
kubectl get all 
```
Можно убедиться, что приложение действительно развертнуто в k8s

## Использование
После запуска приложения используйте браузер или любой http клиент.
Приложение предоставляет методы для просмотра значения переменной `var` в базе данных `Redis`, которое при первом запуске инициализируется как 0.

Доступные методы:
- /add - Прибавить 1 к переменной
- /sub - Отнять 1 от переменной
- /get-current - Вернуть текущее значение переменной
- /save - сохранить состояние базы данных на диск