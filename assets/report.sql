WITH RankedTrades AS (
    SELECT
        T.*,
        ROW_NUMBER() OVER (PARTITION BY T.InstrumentId ORDER BY T.DateEn DESC) AS rn
    FROM
        Trade T
)

SELECT
    I.Name AS instrument_name,
    R.DateEn AS date_en,
    R.Open,
    R.High,
    R.Low,
    R.Close
FROM
    RankedTrades R
        JOIN
    Instrument I ON R.InstrumentId = I.Id
WHERE
    R.rn = 1;
