
function ShowData(){

    let showName = document.getElementById("input-name").value
    let showEmail = document.getElementById("input-email").value
    let showPhone = document.getElementById("input-phone").value
    let showSubject = document.getElementById("input-subject").value
    let showMessage = document.getElementById("input-message").value

    if (showName == "") {
       return alert("name wajib diisi")

    } 
    
    if(showEmail == "") {
       return alert("email wajib diisi")

    }
    
    if(showPhone == "") {
       return alert("phone wajib diisi")

    }
    
    if(showSubject == "") {
       return alert("subject wajib diisi")
       
    }
    
    if(showMessage == "") {
       return alert("message wajib diisi")
    }

    console.log(showName);
    console.log(showEmail);
    console.log(showPhone);
    console.log(showSubject);
    console.log(showMessage);


    let emailReceiver= 'yuzeld432@gmail.com'

    let a = document.createElement('a')
    
    a.href = 'mailto:${emailReceiver}?subject=${showSubject}&body=Hallo nama saya adalah ${showName}, ${showMessage}, silahkan kontak ke nomor ini ${phone}'
   
    a.click()

    

}