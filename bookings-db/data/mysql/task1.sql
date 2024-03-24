-- 1. Выбрать всю информацию о рейсах (flights), в которых номер рейса (flight_no) заканчивается на '488'
SELECT *
FROM flights
WHERE flight_no LIKE '%488'
;

-- 2. Выбрать всю информацию о рейсах (flights), для которых аэропорт Краснодар является пунктом отправления либо прибытия
SELECT *
FROM flights
WHERE departure_airport = 'KRR'
   OR arrival_airport = 'KRR'
;

-- 3. Выбрать всю информацию о рейсах (flights) на самолёте Сухой Суперджет-100, для которых аэропорт Чебоксар
-- является пунктом отправления либо прибытия
SELECT *
FROM flights
WHERE (aircraft_code = 'SU9')
  AND (departure_airport = 'CSY' OR arrival_airport = 'CSY')
;

-- 4. Выбрать идентификаторы и стоимости 10 самых дорогостоящих бронирований (bookings)
SELECT book_ref,
       total_amount
FROM bookings
ORDER BY total_amount DESC
LIMIT 10
;

-- 5. Выбрать имена и контактные данные всех пассажиров, указанных в самом дорогостоящем бронировании
-- (среди всех, что есть в базе данных)
SELECT t.passenger_name,
       t.contact_data
FROM tickets t
         INNER JOIN bookings b ON t.book_ref = b.book_ref
WHERE b.total_amount = (SELECT MAX(total_amount)
                        FROM bookings)
;

-- 6. Выбрать идентификаторы самолётов, в которых есть посадочные места с редким классом 'Comfort'
-- (вместо более привычных 'Economy' / 'Business')
SELECT DISTINCT a.aircraft_code
FROM aircrafts_data a
         INNER JOIN seats s ON a.aircraft_code = s.aircraft_code
WHERE s.fare_conditions = 'Comfort'
;