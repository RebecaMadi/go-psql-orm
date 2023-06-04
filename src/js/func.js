function createDiv(context){
    let div = document.createElement('div');
    context.appendChild(div);
    return div;
}

function createIMG(context){
    let img = document.createElement("IMG");
    context.appendChild(img);
    return img;
}

function createP(context){
    let p = document.createElement('p');
    context.appendChild(p);
    return p;
}

function createForm(context){
    let forms = document.createElement('form');
    context.appendChild(form);
    return form;
}