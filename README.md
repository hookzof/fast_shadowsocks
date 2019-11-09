[English version](README.en.md)

# fast_shadowsocks
Быстрое развёртывание сервера Shadowsocks.

## Работа со скриптом
### Загрузка исполняемого файла
```bash
curl -L -o install https://git.io/JegXf && chmod +x install
```

### Быстрая установка
```bash
./install
```
Генерируется случайный порт от 49152 до 65535 и 10-ти значный пароль.

### Вывод
Выводится два поля **Plain** и **Encoded**.  
В поле **Plain** указаны параметры подключения к серверу - `ss://метод:пароль@хост:порт`  

Для программы <a href="https://getoutline.org/" target="_blank">Outline</a> нужна ссылка из поля **Encoded**.

### Ключ сервера
Сохраняется после установки в файле ss-key.txt, для вывода:
`cat ss-key.txt`

### Файл конфигурации сервера
Путь файла - `/etc/shadowsocks-libev/config.json`  
При изменении необходимо перезагружать сервер командой:  
`service shadowsocks-libev restart`
