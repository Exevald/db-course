-- 1. Для билетов с кодом бронирования '58DF57' выбрать имена пассажиров, номер рейса,
-- дату-время отправления и дату-время прибытия
SELECT DISTINCT t.passenger_name,
       f.flight_no,
       f.scheduled_departure,
       f.scheduled_arrival
FROM tickets t
         INNER JOIN ticket_flights tf ON t.ticket_no = tf.ticket_no
         INNER JOIN flights f ON tf.flight_id = f.flight_id
WHERE t.book_ref = '58DF57'
;

-- 2. Для всех типов самолётов выбрать количество мест по классам обслуживания
SELECT DISTINCT aircraft_code,
                fare_conditions,
                COUNT(aircraft_code)
FROM seats
GROUP BY aircraft_code, fare_conditions
;

-- 3. Выбрать все «счастливые» коды бронирования со списками имён пассажиров в каждом из них
SELECT book_ref,
       GROUP_CONCAT(passenger_name) AS passengers_list
FROM tickets
WHERE SUBSTR(book_ref, 1, 3) = SUBSTR(book_ref, -3)
GROUP BY book_ref
;

-- 4. Выбрать номер рейса, дату-время отправления и дату-время прибытия последнего по времени отправления рейса,
-- прибывшего из Краснодара в Калининград
SELECT flight_no, actual_departure, actual_arrival
FROM flights
WHERE departure_airport = 'KRR'
  AND arrival_airport = 'KGD'
  AND actual_departure IS NOT NULL
  AND actual_arrival IS NOT NULL
ORDER BY actual_departure DESC
LIMIT 1
;

-- 5. Выбрать номер рейса и дату-время отправления для 10 рейсов, принёсших наибольшую выручку
SELECT f.flight_no,
       f.actual_departure,
       SUM(tf.amount) as total_amount
FROM flights f
         INNER JOIN ticket_flights tf ON f.flight_id = tf.flight_id
WHERE status = 'Arrived'
GROUP BY f.flight_no, f.actual_departure
ORDER BY total_amount DESC
LIMIT 10
;