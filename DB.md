## accounts

注意

```sql
CREATE TABLE accounts
(
    account_id uuid NOT NULL DEFAULT gen_random_uuid(),
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    service_type character varying(3) NOT NULL,
    service_id text NOT NULL,
    email character varying(256),
    CONSTRAINT accounts_pkey PRIMARY KEY (account_id, service_type, service_id)
)
```

## boards

```sql
CREATE TABLE public.boards
(
    board_id uuid NOT NULL DEFAULT gen_random_uuid(),
    title character varying(200) NOT NULL,
    body character varying(1000) NOT NULL,
    owner_account_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone,
    CONSTRAINT boards_pkey PRIMARY KEY (board_id)
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE 
    ON public.boards
    FOR EACH ROW
    EXECUTE PROCEDURE public.trigger_set_timestamp();
```

## comments

```sql
CREATE TABLE public.comments
(
    comment_id uuid NOT NULL DEFAULT gen_random_uuid(),
    board_id uuid NOT NULL,
    owner_account_id uuid NOT NULL,
    body character varying(500) NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone,
    CONSTRAINT comments_pkey PRIMARY KEY (comment_id, board_id)
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE 
    ON public.comments
    FOR EACH ROW
    EXECUTE PROCEDURE public.trigger_set_timestamp();
```

## トリガ関数

### trigger_set_timestamp()

参考: https://x-team.com/blog/automatic-timestamps-with-postgresql/

```sql
CREATE FUNCTION public.trigger_set_timestamp()
    RETURNS trigger
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE NOT LEAKPROOF
AS $BODY$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$BODY$;
```


