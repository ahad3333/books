
CREATE TABLE books (
    id UUID  PRIMARY KEY,
    name VARCHAR NOT NULL,
    price NUMERIC NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
CREATE TABLE categorys (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL
);
CREATE TABLE BookCategory (
    id UUID PRIMARY KEY,
    bookId UUID REFERENCES books(id),
    categoryId UUID REFERENCES categorys(id)
);

insert into BookCategory(id,bookId,categoryId)
VALUES 
('9ff8fcd0-660c-4d2d-a5d8-16a676e47818','baf34639-ed47-4ee3-a050-6398d42ac06d','34d72a54-36e3-40eb-8944-078f82dfe12b'),
('50a0101c-2098-4df2-97ea-dd77d63d40c4','9da542d0-a54e-44f6-900f-5250e866b747','52079474-a2e4-4d7b-9951-e146ca2b2e10'),
('729f2dac-8ea5-11ed-a1eb-0242ac120002','9da542d0-a54e-44f6-900f-5250e866b747','f9fd675a-4c11-4ea6-8872-f33efee2e122'),
('7c72e5a8-8ea5-11ed-a1eb-0242ac120002','8f85e52d-ebb1-44ea-ae20-331034280a23','be936be5-609d-4370-80b0-fb1b106857b0'),
('84ff6390-8ea5-11ed-a1eb-0242ac120002','8f85e52d-ebb1-44ea-ae20-331034280a23','e04e7ac0-db2d-4371-ab33-cc5776d896d0');

select
b.name as name,
b.price as price,
b.description as description,
array_agg(c.name) 
from BookCategory as bc
join books as b on b.id = bc.bookId
join categorys as c on c.id = bc.categoryId
where b.id = '9da542d0-a54e-44f6-900f-5250e866b747'
group by b.name,b.price, b.description;

select
count(*) OVER(),
	b.name,
	b.price,
	b.description,
	arrey_agg(c.name)
	from books as b
	join BookCategory as bc on bc.bookId = b.id
	join categorys as c on c.id = bc.categoryId
	where b.id = '9da542d0-a54e-44f6-900f-5250e866b747'
	group by b.name,b.price, b.description,c.name;

select
array_agg(c.name)
from categorys as c
 join BookCategory as bc on bc.categoryId = c.id 
 join books as b on b.id = bc.bookId
 where b.id = '9da542d0-a54e-44f6-900f-5250e866b747'
 group by c.name;  

select
array_agg(c.name)
from BookCategory as cb
join categorys as c on c.id = cb.categoryId
where cb.bookId = '9da542d0-a54e-44f6-900f-5250e866b747';
group by c.name;

select
b.name
from BookCategory as cb
join books as b on b.id = cb.bookId
where b.id = '9da542d0-a54e-44f6-900f-5250e866b747'
group by b.name;