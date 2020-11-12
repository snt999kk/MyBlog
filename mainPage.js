let name = document.getElementById('1');
let main = document.getElementById('2');
name.onclick = function()
{
    let val = document.getElementById('3').value;
    if(val != '')
        main.textContent = 'Привет ' + val;
}