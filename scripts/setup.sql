CREATE TABLE parameters (
        Id SERIAL PRIMARY KEY,
        CreateDate timestamp,
        K INT,
        M INT,
        N INT
);

CREATE TABLE addressLists (
        Id SERIAL PRIMARY KEY,
        CreateDate timestamp,
    	IsWhite BOOLEAN DEFAULT false,
	Network VARCHAR (20) UNIQUE
);


INSERT INTO parameters (CreateDate, N, M, K) VALUES ('2019-10-19 19:00:00', 3, 6 ,9);
INSERT INTO addressLists (CreateDate, IsWhite, Network) VALUES ('2019-10-19 19:00:00', true, '127.0.0.1/24');
INSERT INTO addressLists (CreateDate, IsWhite, Network) VALUES ('2019-10-19 19:00:00', false, '127.0.1.1/24');