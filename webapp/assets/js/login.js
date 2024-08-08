$('#login').on('submit', fazerLogin);

function fazerLogin(evento) {
  evento.preventDefault();
  $.ajax({
    url: '/login',
    method: 'POST',
    data: {
      email: $('#email').val(),
      senha: $('#senha').val()
    }
  }).done(function(){
    Swal.fire(
      'Sucesso!',
      'Usuário logado com sucesso!',
      'success'
    ).then(function(){
      window.location = '/home'
    })
  }).fail(function (erro) {
     Swal.fire("Ops...", "Usuário ou senha inválidos!", "error")
  })
}