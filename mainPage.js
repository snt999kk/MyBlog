let name = document.getElementById('1');
let main = document.getElementById('2');
name.onclick = function()
{
    main.textContent = 'Привет ' + document.getElementById('3').value
}