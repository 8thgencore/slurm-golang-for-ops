## Решите задачу для закрепления навыков по материалам видеолекции:

Для того, чтобы освоить основы пакетов `net/http` и `os`, напишите http сервер с методом `POST` `/log`,который принимает строку и помещает её в новую линию в файле, путь к которому определен переменной env `APP_LOGFILE_PATH`. 

Если путь не определен, то сервер должен использовать свой корень, где он запущен с именем файла `log.txt` по умолчанию. После записи строки в лог сервер должен вернуть код `200` и `OK`