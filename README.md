O Korvia é uma ferramenta, que realiza agendamentos de serviços através do whatsapp, ele tem o seguinte fluxo:

Cliente
↓

Menu

↓

Escolher serviço

↓

Como deseja agendar?

↓

Primeiro horário
│
└────► Confirmação

ou

Escolher horário
│
├── Próximo dia
├── Explorar agenda
└── Horário escolhido
│
▼
Escolher profissional
│
▼
Confirmação

Cliente

- possui um nome
- possui um contato
- possui um email
- possui uma data de nascimento

Cadastrar()

Alterar()

Professional

- possui um nome
- possui um contato
- possui uma função (owner, manager, receptionist, professional)

Cadastrar()

Alterar()

Deletar()

Serviço

- possui tempo de duração
- possui um valor R$

Cadastrar()

Alterar()

Deletar()

Agendamento

- pertence a uma Barbearia
- possui um Cliente
- possui um Profissional
- possui um Serviço
- possui data
- possui horário
- possui duração
- possui status

Confirmar()

Cancelar()

Reagendar()

Concluir()
