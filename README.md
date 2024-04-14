# wp24-athletes

This is a simple open REST API for immaginary ancient greece athletes.
It was written for the third mini homework of the
[Web Programming M-Z](https://perceivelab.github.io/wp-mz-24) @ DIEEI
course.

## Endpoints

### /

Return information about the API.

For example, `curl -s http://<URL of the athletes API>/ | jq`:

```json
{
  "description": "A simple open REST API for athletes!",
  "version": "0.1"
}
```

### /athletes

Return the alphabetically ordered list of athletes.

For example, `curl -s http://<URL of the athletes API>/athletes | jq`:

```
[
  {
    "slug": "arrhichione-di-fliunte",
    "name": "Arrhichione di Fliunte",
    "gender": "M",
    "age": 26,
    "sport": "Lotta",
    "famous_for": "Conosciuto per il suo successo nella lotta e il suo sacrificio durante una competizione."
  },
  {
    "slug": "astioco-di-taranto",
    "name": "Astioco di Taranto",
    "gender": "M",
    "age": 25,
    "sport": "Lotta",
    "famous_for": "Conosciuto per la sua forza e abilità nella lotta."
  },
  {
    "slug": "callias-di-atene",
    "name": "Callias di Atene",
    "gender": "M",
    "age": 24,
    "sport": "Danza",
    "famous_for": "Rinomato per la sua destrezza nella danza."
  },
  {
    "slug": "chryseis-di-larissa",
    "name": "Chryseis di Larissa",
    "gender": "F",
    "age": 24,
    "sport": "Lancio del giavellotto",
    "famous_for": "Famosa per la sua agilità e velocità nel lancio del giavellotto."
  },

  MORE ENTRIES...

]
```
