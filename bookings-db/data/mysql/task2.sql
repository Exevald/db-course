-- 1. Для всех рейсов Домодедово, находящихся в статусе 'Delayed', поменять статус на 'Cancelled'
UPDATE flights
SET status = 'Cancelled'
WHERE (departure_airport = 'DME' OR arrival_airport = 'DME')
  AND status = 'Delayed'
;

-- Обратный запрос
UPDATE flights
SET status = 'Delayed'
WHERE (departure_airport = 'DME' OR arrival_airport = 'DME')
  AND status = 'Cancelled'
  AND (actual_departure IS NULL AND actual_arrival IS NULL)
;

SELECT *
FROM flights
WHERE status = 'Cancelled'
  AND (
    departure_airport = 'DME'
        OR arrival_airport = 'DME')
;

-- 2. Для всех рейсов аэропорта Йошкар-Олы, находящихся в статусе 'Scheduled', поменять статус на 'Arrived'
-- и установить фактические даты вылета и прилёта равными запланированным
SELECT *
FROM flights
WHERE (departure_airport = 'JOK' OR arrival_airport = 'JOK')
  AND status = 'Arrived'
  AND actual_departure = scheduled_departure
  AND actual_arrival = scheduled_arrival
;

UPDATE flights
SET status           = 'Arrived',
    actual_departure = scheduled_departure,
    actual_arrival   = scheduled_arrival
WHERE (departure_airport = 'JOK' OR arrival_airport = 'JOK')
  AND status = 'Scheduled'
;

-- Обратный запрос
UPDATE flights
SET status           = 'Scheduled',
    actual_departure = NULL,
    actual_arrival   = NULL
WHERE (departure_airport = 'JOK' OR arrival_airport = 'JOK')
  AND status = 'Arrived'
  AND actual_departure = scheduled_departure
  AND actual_arrival = scheduled_arrival
;

UPDATE flights
SET status           = 'Arrived',
    actual_departure = scheduled_departure,
    actual_arrival   = scheduled_arrival
WHERE flight_id IN (1997, 33347, 33422, 33624, 33642, 47884)
;

SELECT COUNT(*)
FROM flights
WHERE status = 'Scheduled'
  AND (departure_airport = 'JOK' OR arrival_airport = 'JOK')
;

-- 3. Удалить всю информацию о билетах пассажира Gennadiy Nikitin
SELECT *
FROM ticket_flights
WHERE ticket_no IN (SELECT ticket_no
                    FROM tickets
                    WHERE passenger_name = 'Gennadiy Nikitin')
;

SELECT *
FROM boarding_passes
WHERE ticket_no IN (SELECT ticket_no
                    FROM tickets
                    WHERE passenger_name = 'Gennadiy Nikitin')
;

SELECT *
FROM bookings
WHERE book_ref IN (SELECT book_ref
                   FROM tickets
                   WHERE passenger_name = 'Gennadiy Nikitin')
;

SELECT *
FROM tickets
WHERE passenger_name = 'Gennadiy Nikitin'
;

DELETE
FROM ticket_flights
WHERE ticket_no IN (SELECT ticket_no
                    FROM tickets
                    WHERE passenger_name = 'Gennadiy Nikitin')
;

DELETE
FROM boarding_passes
WHERE ticket_no IN (SELECT ticket_no
                    FROM tickets
                    WHERE passenger_name = 'Gennadiy Nikitin')
;

DELETE
FROM bookings
WHERE book_ref IN (SELECT book_ref
                   FROM tickets
                   WHERE passenger_name = 'Gennadiy Nikitin')
;

DELETE
FROM tickets
WHERE passenger_name = 'Gennadiy Nikitin'
;

-- Обратный запрос