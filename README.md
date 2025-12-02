# Практическое задание № 13 Борисов Денис Александрович ЭФМО-01-25

Цели:
1.	Научиться подключать и использовать профилировщик pprof для анализа CPU, памяти, блокировок и горутин.
2.	Освоить базовые техники измерения времени выполнения функций (ручные таймеры, бенчмарки).
3.	Научиться читать отчёты go tool pprof, строить графы вызовов и находить “узкие места”.
4.	Сформировать практические навыки оптимизации кода на основании метрик.

# Выполнение практического задания

1. Шаблон проекта

Для выполнения практической работы была сделана следующая структура проекта

<img width="179" height="225" alt="image" src="https://github.com/user-attachments/assets/9bb8289a-e1bf-42a0-a47c-bde235f850a8" />


2. HTTP-сервер с pprof

Для выполнения замеров нагрузки в рамках практики были сделаны файлы slow.go и main.go

Содержание файла cmd/api/main.go

<img width="641" height="513" alt="image" src="https://github.com/user-attachments/assets/45939df9-a2ce-48c2-9565-7ed070160448" />

Содержание файла internal/work/slow.go

<img width="303" height="396" alt="image" src="https://github.com/user-attachments/assets/7a89797d-6819-4523-a817-0daf99b96632" />

3. Получение профилей через HTTP

Посел был запущен сервер и были получены следую.щие данные:

*Индекс pprof

<img width="1693" height="765" alt="image" src="https://github.com/user-attachments/assets/17585944-f67b-4dd3-bcdf-59133db36069" />

4. Анализ профилей

5. Ручное измерение времени функций

Для ручного измерения времени был создан файл timer.go

<img width="505" height="307" alt="image" src="https://github.com/user-attachments/assets/9ab4a70e-ad55-405a-955d-976747b79809" />

6. Бенчмарки: стабильные метрики

Создайте тест с бенчмарком internal/work/slow_test.go

<img width="504" height="365" alt="image" src="https://github.com/user-attachments/assets/9617e717-aa0f-4e82-aba7-e0a20a65eb46" />

Результат выполнения

<img width="632" height="238" alt="image" src="https://github.com/user-attachments/assets/e990426d-3c53-49c7-9017-39473cf4cb37" />

7. Проба оптимизации и повторные замеры

После был изменен файл slow.go добавление функции FibFast 

<img width="307" height="398" alt="image" src="https://github.com/user-attachments/assets/7c92b21b-3341-4057-8bdf-14ce997cf8c9" />

И был сделана вторая Бенчмарка

<img width="372" height="138" alt="image" src="https://github.com/user-attachments/assets/0f450849-0c9e-4bb5-aeb5-12b167a263b7" />

