# Programming Paradigms

## Functional

Created in 1957

## Procedural

Created in 1968

## Object Oriented

Created in 1980/90

## Imperative

- Foco no fluxo
- Estados mutáveis
- Como
- Maior quantidade de código
- Baixo nível de escalabilidade
- Mais conhecido
- Mais explícito
  Example

```javascript
const notas = [8.7, 6.8, 7.7, 7.7, 9.2, 5.3, 8.0];

function media(notas) {
  let total = 0;
  for (let i = 0; i < notas.length; i++) {
    total += notas[i];
  }

  return total / notas.length;
}

const mediaTurma = media(notas);
console.log(`Média é ${mediaTurma}`);
```

## Declarative

- Foco na lógica
- Imutabilidade
- O que
- Menor quantidade de código
- Alto nível de escalabilidade
- Menos conhecido
- Menos explícito
  Example

```SQL
SELECT
    COUNT(*) monthAccesses
FROM user_accesses
WHERE year(curdate()) = year(last_activity_at)
AND month(curdate()) = month(last_activity_at);
```

```javascript
const notas = [8.7, 6.8, 7.7, 7.7, 9.2, 5.3, 8.0];

const somar = (a, b) => a + b;
const dividir = (a, b) => a / b;

function media(notas) {
  let total = 0;
  for (let i = 0; i < notas.length; i++) {
    total += notas[i];
  }

  return total / notas.length;
}

const mediaTurma = media(notas);
console.log(`Média é ${mediaTurma}`);

const mediaTurma = dividir(notas.reduce(somar), notas.length);

console.log(`Média é ${mediaTurma}`);
```
