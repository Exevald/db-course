-- 1. Для всех рейсов Домодедово, находящихся в статусе 'Delayed', поменять статус на 'Cancelled'
UPDATE flights
SET status = 'Cancelled'
WHERE (departure_airport = 'DME' OR arrival_airport = 'DME')
  AND status = 'Delayed'
;

-- Обратный запрос
UPDATE flights
SET status = 'Delayed'
WHERE flight_id IN (348, 761, 974, 2469, 5377, 5858, 34275, 36780, 46165, 53170, 59329)
;

-- 2. Для всех рейсов аэропорта Йошкар-Олы, находящихся в статусе 'Scheduled', поменять статус на 'Arrived'
-- и установить фактические даты вылета и прилёта равными запланированным

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
WHERE flight_id IN
      (1937, 1940, 1941, 1942, 1944, 1949, 1952, 1961, 1970, 1971, 1973, 1974, 1976, 1978, 1980, 1996, 2008, 2010, 2011,
       2012, 2025, 2026, 2033, 2039, 2042, 2046, 2052, 2053, 7351, 7352, 7364, 7365, 7366, 7367, 7371, 7374, 7387, 7388,
       7401, 7402, 7405, 7406, 7411, 7412, 7415, 7416, 7426, 7427, 7428, 7429, 7442, 7450, 7451, 7458, 7479, 7481, 7482,
       7484, 7494, 7501, 7503, 7505, 7513, 7514, 7515, 7516, 7520, 7528, 7529, 7531, 7532, 7533, 7534, 7535, 7537,
       7538, 7539, 7542, 7562, 7569, 7570, 7571, 7572, 7580, 7583, 11597, 11599, 11601, 11604, 11607, 11616, 11617,
       11622, 11623, 11629, 11641, 11644, 11645, 11657, 11663, 11664, 11678, 11680, 11681, 11683, 11684, 11690, 11697,
       11703, 11704, 11710, 11713, 11714, 33193, 33195, 33200, 33201, 33204, 33205, 33206, 33208, 33225, 33229, 33233,
       33238, 33242, 33244, 33246, 33249, 33250, 33254, 33259, 33262, 33264, 33265, 33267, 33269, 33288, 33291, 33299,
       33304, 33307, 33314, 33315, 33324, 33325, 33328, 33331, 33332, 33333, 33338, 33343, 33344, 33345, 33346, 33361,
       33362, 33369, 33370, 33378, 33379, 33382, 33383, 33395, 33398, 33405, 33406, 33420, 33421, 33425, 33427, 33428,
       33429, 33438, 33439, 33444, 33449, 33454, 33455, 33458, 33462, 33463, 33508, 33518, 33519, 33520, 33521, 33522,
       33523, 33524, 33525, 33527, 33529, 33530, 33534, 33535, 33542, 33543, 33549, 33554, 33564, 33571, 33577, 33578,
       33586, 33591, 33602, 33603, 33606, 33609, 33613, 33614, 33618, 33625, 33626, 33627, 33631, 33634, 33636, 33639,
       33649, 33652, 33658, 33660, 33663, 33670, 33674, 33677, 33684, 33689, 33690, 33696, 33701, 33705, 33706, 47858,
       47860, 47873, 47876, 47886, 47887, 47889
          )
;