/*const container = document.getElementById('container');
const loginButton = document.getElementById('login');
const signUpButton = document.getElementById('signUp');

signUpButton.addEventListener('click',() => {
    container.classList.add('panel-active');
})
loginButton.addEventListener('click',() => {
    container.classList.remove('panel-active');
})*/
const container = document.getElementById('container');
const signUpButton = document.getElementById('signUp');  // Utilisez le même identifiant que dans votre bouton HTML
const loginButton = document.getElementById('login');     // Utilisez le même identifiant que dans votre bouton HTML

signUpButton.addEventListener('click', () => {
    container.classList.add('panel-active');
});

loginButton.addEventListener('click', () => {
    container.classList.remove('panel-active');
});