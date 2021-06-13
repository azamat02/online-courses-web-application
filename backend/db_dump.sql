--
-- PostgreSQL database dump
--

-- Dumped from database version 13.2
-- Dumped by pg_dump version 13.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: comments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.comments (
    id bigint NOT NULL,
    user_id bigint,
    course_id bigint,
    ctext text,
    rate bigint,
    created_date timestamp with time zone
);


ALTER TABLE public.comments OWNER TO postgres;

--
-- Name: comments_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.comments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.comments_id_seq OWNER TO postgres;

--
-- Name: comments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.comments_id_seq OWNED BY public.comments.id;


--
-- Name: completed_lesson_logs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.completed_lesson_logs (
    id bigint NOT NULL,
    user_id bigint,
    course_id bigint,
    lesson_id bigint
);


ALTER TABLE public.completed_lesson_logs OWNER TO postgres;

--
-- Name: completed_lesson_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.completed_lesson_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.completed_lesson_logs_id_seq OWNER TO postgres;

--
-- Name: completed_lesson_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.completed_lesson_logs_id_seq OWNED BY public.completed_lesson_logs.id;


--
-- Name: course_analytics_logs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.course_analytics_logs (
    id bigint NOT NULL,
    user_id bigint,
    course_id bigint,
    log text,
    date timestamp with time zone
);


ALTER TABLE public.course_analytics_logs OWNER TO postgres;

--
-- Name: course_analytics_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.course_analytics_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.course_analytics_logs_id_seq OWNER TO postgres;

--
-- Name: course_analytics_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.course_analytics_logs_id_seq OWNED BY public.course_analytics_logs.id;


--
-- Name: course_progresses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.course_progresses (
    id bigint NOT NULL,
    course_id bigint,
    user_id bigint,
    passed_lessons_count bigint
);


ALTER TABLE public.course_progresses OWNER TO postgres;

--
-- Name: course_progresses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.course_progresses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.course_progresses_id_seq OWNER TO postgres;

--
-- Name: course_progresses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.course_progresses_id_seq OWNED BY public.course_progresses.id;


--
-- Name: courses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.courses (
    id bigint NOT NULL,
    img text,
    title text,
    description text,
    created_data timestamp with time zone,
    req text,
    what_you_will_learn text,
    category text
);


ALTER TABLE public.courses OWNER TO postgres;

--
-- Name: courses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.courses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.courses_id_seq OWNER TO postgres;

--
-- Name: courses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.courses_id_seq OWNED BY public.courses.id;


--
-- Name: lessons; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.lessons (
    id bigint NOT NULL,
    type text,
    module_id bigint,
    title text,
    link text,
    content text
);


ALTER TABLE public.lessons OWNER TO postgres;

--
-- Name: lessons_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.lessons_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lessons_id_seq OWNER TO postgres;

--
-- Name: lessons_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.lessons_id_seq OWNED BY public.lessons.id;


--
-- Name: log_of_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.log_of_users (
    id bigint NOT NULL,
    user_id bigint,
    enter_date timestamp with time zone
);


ALTER TABLE public.log_of_users OWNER TO postgres;

--
-- Name: log_of_users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.log_of_users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.log_of_users_id_seq OWNER TO postgres;

--
-- Name: log_of_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.log_of_users_id_seq OWNED BY public.log_of_users.id;


--
-- Name: modules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.modules (
    id bigint NOT NULL,
    title text,
    course_id bigint,
    number_of_lessons bigint
);


ALTER TABLE public.modules OWNER TO postgres;

--
-- Name: modules_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.modules_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.modules_id_seq OWNER TO postgres;

--
-- Name: modules_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.modules_id_seq OWNED BY public.modules.id;


--
-- Name: purchased_courses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.purchased_courses (
    id bigint NOT NULL,
    user_id bigint,
    course_id bigint,
    purchased_date timestamp with time zone
);


ALTER TABLE public.purchased_courses OWNER TO postgres;

--
-- Name: purchased_courses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.purchased_courses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.purchased_courses_id_seq OWNER TO postgres;

--
-- Name: purchased_courses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.purchased_courses_id_seq OWNED BY public.purchased_courses.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    login text,
    password bytea,
    name text,
    surname text,
    email text
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: comments id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments ALTER COLUMN id SET DEFAULT nextval('public.comments_id_seq'::regclass);


--
-- Name: completed_lesson_logs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.completed_lesson_logs ALTER COLUMN id SET DEFAULT nextval('public.completed_lesson_logs_id_seq'::regclass);


--
-- Name: course_analytics_logs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course_analytics_logs ALTER COLUMN id SET DEFAULT nextval('public.course_analytics_logs_id_seq'::regclass);


--
-- Name: course_progresses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course_progresses ALTER COLUMN id SET DEFAULT nextval('public.course_progresses_id_seq'::regclass);


--
-- Name: courses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.courses ALTER COLUMN id SET DEFAULT nextval('public.courses_id_seq'::regclass);


--
-- Name: lessons id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lessons ALTER COLUMN id SET DEFAULT nextval('public.lessons_id_seq'::regclass);


--
-- Name: log_of_users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.log_of_users ALTER COLUMN id SET DEFAULT nextval('public.log_of_users_id_seq'::regclass);


--
-- Name: modules id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.modules ALTER COLUMN id SET DEFAULT nextval('public.modules_id_seq'::regclass);


--
-- Name: purchased_courses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.purchased_courses ALTER COLUMN id SET DEFAULT nextval('public.purchased_courses_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: comments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.comments (id, user_id, course_id, ctext, rate, created_date) FROM stdin;
13	1	1	good	5	2021-06-13 21:00:03.793818+03
\.


--
-- Data for Name: completed_lesson_logs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.completed_lesson_logs (id, user_id, course_id, lesson_id) FROM stdin;
14	1	1	1
15	1	1	2
16	1	1	3
17	1	1	6
\.


--
-- Data for Name: course_analytics_logs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.course_analytics_logs (id, user_id, course_id, log, date) FROM stdin;
42	1	1	Entered to lesson	2021-05-28 09:19:18.936824+03
43	1	1	Entered to lesson	2021-05-28 09:20:23.049785+03
44	1	1	Completed lesson	2021-05-28 09:20:23.054065+03
45	1	1	Entered to lesson	2021-05-28 09:21:00.624592+03
46	1	1	Completed lesson	2021-05-28 09:21:00.624592+03
47	1	1	Entered to lesson	2021-05-28 09:21:07.241335+03
48	1	1	Entered to lesson	2021-05-28 09:24:42.422037+03
49	1	1	Entered to lesson	2021-05-28 09:24:47.105072+03
50	1	1	Completed lesson	2021-05-28 09:24:49.619599+03
51	1	1	Entered to lesson	2021-05-28 09:24:55.392465+03
52	1	1	Entered to lesson	2021-06-01 11:03:16.992133+03
53	1	1	Entered to lesson	2021-06-02 10:53:57.722498+03
54	1	1	Entered to lesson	2021-06-02 11:03:11.407871+03
55	1	1	Entered to lesson	2021-06-13 11:57:01.989612+03
56	1	1	Entered to lesson	2021-06-13 11:57:13.388492+03
57	1	1	Entered to lesson	2021-06-13 17:22:41.746648+03
58	1	1	Entered to lesson	2021-06-13 17:22:55.25269+03
59	1	1	Entered to lesson	2021-06-13 17:27:50.548962+03
60	1	1	Entered to lesson	2021-06-13 17:28:24.523194+03
61	1	1	Entered to lesson	2021-06-13 17:28:55.399111+03
62	1	1	Entered to lesson	2021-06-13 17:29:08.229691+03
63	1	1	Entered to lesson	2021-06-13 17:29:57.83724+03
64	1	1	Entered to lesson	2021-06-13 17:32:27.071421+03
65	10	1	Entered to lesson	2021-06-13 17:54:34.183066+03
66	10	1	Entered to lesson	2021-06-13 20:47:18.846331+03
67	1	1	Entered to lesson	2021-06-13 20:47:37.118815+03
68	1	1	Entered to lesson	2021-06-13 20:47:46.866773+03
69	1	1	Entered to lesson	2021-06-13 20:52:28.796515+03
70	1	1	Entered to lesson	2021-06-13 20:53:05.328886+03
71	1	1	Entered to lesson	2021-06-13 20:53:33.229699+03
72	1	1	Entered to lesson	2021-06-13 20:53:35.731048+03
73	1	1	Entered to lesson	2021-06-13 20:53:58.093458+03
74	1	1	Entered to lesson	2021-06-13 20:54:39.3872+03
75	1	1	Entered to lesson	2021-06-13 20:54:47.232828+03
76	1	1	Completed lesson	2021-06-13 20:54:53.725233+03
77	1	1	Entered to lesson	2021-06-13 20:54:58.489394+03
78	1	1	Entered to lesson	2021-06-13 20:58:38.639046+03
79	10	1	Entered to lesson	2021-06-13 21:02:44.108511+03
80	1	1	Entered to lesson	2021-06-13 21:09:17.081359+03
81	1	1	Entered to lesson	2021-06-13 21:14:35.427643+03
\.


--
-- Data for Name: course_progresses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.course_progresses (id, course_id, user_id, passed_lessons_count) FROM stdin;
\.


--
-- Data for Name: courses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.courses (id, img, title, description, created_data, req, what_you_will_learn, category) FROM stdin;
2	https://bit.ly/3tDioR8	Golang course	Golang crash course for beginners	2021-05-16 21:05:09.6362+03	HTML/CSS/JS;OOP;Git;Bash;Computer	\tGolang language; OOP concepts in golang; Databases(MySQL, SQLite3, PostgreSQL); Building web application on golang;Simple to understand	Programming
1	https://bit.ly/3heqU6H	Python course	Python crash course	2021-05-16 20:47:07.629188+03	HTML/CSS/JS;OOP;Git;Bash	Python 3.8; OOP concepts in python; Databases(MySQL, SQLite3)	Programming
\.


--
-- Data for Name: lessons; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.lessons (id, type, module_id, title, link, content) FROM stdin;
2	video	1	Installing python 3.8	https://www.youtube.com/watch?v=UvcQlPZ8ecA	Today we will install python 3.8
3	lecture	1	Setting python in system path		The PATH variable is a list of directories where each directory contains the executable file for a command.When a command is entered into the Windows command prompt, the prompt searches in the PATH variable for an executable file with the same name as the command; in the case that the required file is not found, it responds with an error message that states that the specified command was not recognized.
1	video	1	Install PyCharm	https://www.youtube.com/watch?v=SZUNUB6nz3g	Some content here
6	video	1	Test	https://www.a2hosting.com/kb/developer-corner/postgresql/import-and-export-a-postgresql-database#:~:text=In%20the%20left%20pane%20of,Format%20list%20box%2C%20select%20SQL.	asasas
\.


--
-- Data for Name: log_of_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.log_of_users (id, user_id, enter_date) FROM stdin;
1	1	2021-05-18 16:14:32.824271+03
2	1	2021-05-18 16:20:40.400575+03
3	1	2021-05-18 17:19:54.987079+03
5	1	2021-05-19 08:04:45.803757+03
7	1	2021-05-19 17:56:05.605102+03
8	1	2021-05-20 08:15:05.147036+03
12	1	2021-05-20 09:03:31.872509+03
14	1	2021-05-20 09:14:19.986863+03
15	4	2021-05-20 09:15:11.571711+03
16	1	2021-05-20 09:24:26.275023+03
17	1	2021-05-20 09:33:44.491299+03
18	1	2021-05-20 09:39:09.922095+03
19	4	2021-05-20 09:41:43.296567+03
20	1	2021-05-20 09:42:01.901472+03
21	1	2021-05-24 09:42:40.279713+03
22	1	2021-05-24 17:06:45.809937+03
23	5	2021-05-24 17:07:05.388044+03
24	1	2021-05-24 17:08:02.427916+03
25	1	2021-05-26 08:44:26.589873+03
26	1	2021-05-27 12:23:17.487319+03
27	1	2021-05-27 12:24:24.127607+03
28	1	2021-05-31 08:32:50.539534+03
29	1	2021-06-01 08:08:46.408155+03
30	1	2021-06-02 10:02:54.848696+03
31	1	2021-06-03 06:46:10.01999+03
32	1	2021-06-03 06:57:16.993755+03
33	5	2021-06-03 06:58:07.933744+03
34	1	2021-06-03 07:13:54.791386+03
35	5	2021-06-03 07:14:30.939846+03
36	1	2021-06-13 09:36:02.456076+03
37	10	2021-06-13 17:54:20.306348+03
38	1	2021-06-13 20:47:32.767353+03
39	10	2021-06-13 21:02:27.523783+03
40	1	2021-06-13 21:08:34.510097+03
41	1	2021-06-13 21:13:43.25868+03
\.


--
-- Data for Name: modules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.modules (id, title, course_id, number_of_lessons) FROM stdin;
1	Module 1, setting up environment	1	7
2	Module 2, creating first project	1	10
3	Module 3, additional information	1	5
4	Module 1, setting up environment	2	12
5	Module 2, creating first project	2	2
6	\tModule 3, additional information	2	12
\.


--
-- Data for Name: purchased_courses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.purchased_courses (id, user_id, course_id, purchased_date) FROM stdin;
1	1	1	2021-05-18 16:36:09.046574+03
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, login, password, name, surname, email) FROM stdin;
1	aza123	\\x243261243134244553303678376552334367596634376d487231716f6547364b4576334e77704d7336662f685a6d54735663765952596e354c616a6d	Азамат	Саидулы	azamattolegenov1@gmail.com
5	jhon123	\\x2432612431342437646b3577655a6a2f43437056626578395a70647175626a73397a5a534a4d4f7842754d6a49335553327951556e4973723279644b	Jhon	Doe	jhon@test.com
6	kaliev	\\x243261243134246b4737336d716573517355335836317a38516e6e57757452714e7659522e71345738773151426e726d42754948464d492f55452f43	Aniyar	Kaliev	kaliev@mail.com
7	nurs22	\\x243261243134244d62746d5764513947464a30536c55543533512f6a65625043472e347134672f6f6e6e476278584b4a6431776374486550334e4671	Nursultan	Yegenberdiuly	nurs@mail.com
4	mark234	\\x243261243134244b426554654664647a55654532623753394b4c36662e3278336c6d4c6a37364868642f654e556b5646367368756d2f71776b777a75	Mark	Doe	mark@mail.com
9	test123	\\x243261243134242e5176587761724b3170374a595a533131797033567566595947576133616a395973447343594c707878616f6441734f7a4d68674b	test	test	test@mail.com
10	mark123	\\x243261243134246e41697152766a356a6268722e6d4f507a366159302e684e62704a4a52434e316275683871636b446533644c5779444930487a2e32	Mark	Doe	mark1@mail.com
\.


--
-- Name: comments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.comments_id_seq', 13, true);


--
-- Name: completed_lesson_logs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.completed_lesson_logs_id_seq', 17, true);


--
-- Name: course_analytics_logs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.course_analytics_logs_id_seq', 81, true);


--
-- Name: course_progresses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.course_progresses_id_seq', 1, false);


--
-- Name: courses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.courses_id_seq', 6, true);


--
-- Name: lessons_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.lessons_id_seq', 6, true);


--
-- Name: log_of_users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.log_of_users_id_seq', 41, true);


--
-- Name: modules_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.modules_id_seq', 7, true);


--
-- Name: purchased_courses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.purchased_courses_id_seq', 1, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 10, true);


--
-- Name: comments comments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_pkey PRIMARY KEY (id);


--
-- Name: completed_lesson_logs completed_lesson_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.completed_lesson_logs
    ADD CONSTRAINT completed_lesson_logs_pkey PRIMARY KEY (id);


--
-- Name: course_analytics_logs course_analytics_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course_analytics_logs
    ADD CONSTRAINT course_analytics_logs_pkey PRIMARY KEY (id);


--
-- Name: course_progresses course_progresses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course_progresses
    ADD CONSTRAINT course_progresses_pkey PRIMARY KEY (id);


--
-- Name: courses courses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT courses_pkey PRIMARY KEY (id);


--
-- Name: lessons lessons_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lessons
    ADD CONSTRAINT lessons_pkey PRIMARY KEY (id);


--
-- Name: log_of_users log_of_users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.log_of_users
    ADD CONSTRAINT log_of_users_pkey PRIMARY KEY (id);


--
-- Name: modules modules_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.modules
    ADD CONSTRAINT modules_pkey PRIMARY KEY (id);


--
-- Name: purchased_courses purchased_courses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.purchased_courses
    ADD CONSTRAINT purchased_courses_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: comments fk_comments_course; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT fk_comments_course FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- Name: comments fk_comments_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT fk_comments_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: completed_lesson_logs fk_completed_lesson_logs_course; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.completed_lesson_logs
    ADD CONSTRAINT fk_completed_lesson_logs_course FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- Name: completed_lesson_logs fk_completed_lesson_logs_lesson; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.completed_lesson_logs
    ADD CONSTRAINT fk_completed_lesson_logs_lesson FOREIGN KEY (lesson_id) REFERENCES public.lessons(id);


--
-- Name: completed_lesson_logs fk_completed_lesson_logs_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.completed_lesson_logs
    ADD CONSTRAINT fk_completed_lesson_logs_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: course_analytics_logs fk_course_analytics_logs_course; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course_analytics_logs
    ADD CONSTRAINT fk_course_analytics_logs_course FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- Name: course_analytics_logs fk_course_analytics_logs_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course_analytics_logs
    ADD CONSTRAINT fk_course_analytics_logs_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: course_progresses fk_course_progresses_course; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course_progresses
    ADD CONSTRAINT fk_course_progresses_course FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- Name: course_progresses fk_course_progresses_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course_progresses
    ADD CONSTRAINT fk_course_progresses_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: lessons fk_lessons_module; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lessons
    ADD CONSTRAINT fk_lessons_module FOREIGN KEY (module_id) REFERENCES public.modules(id);


--
-- Name: log_of_users fk_log_of_users_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.log_of_users
    ADD CONSTRAINT fk_log_of_users_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: modules fk_modules_course; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.modules
    ADD CONSTRAINT fk_modules_course FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- Name: purchased_courses fk_purchased_courses_course; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.purchased_courses
    ADD CONSTRAINT fk_purchased_courses_course FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- Name: purchased_courses fk_purchased_courses_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.purchased_courses
    ADD CONSTRAINT fk_purchased_courses_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

