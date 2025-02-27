function Range() {
    const slider = document.getElementById("MembersNumber");
    const val = slider.value;
    if (val>0) {
        document.getElementById("LabelMembers").innerHTML = val
    } else {
        document.getElementById("LabelMembers").innerHTML = "All"
    }

}

function myFunction() {
    const input = document.getElementById("searchBar");
    const filter = input.value.toUpperCase();
    const ul = document.getElementById("groups");
    const li = ul.getElementsByTagName("li");

    for (i = 0; i < li.length; i++) {
        const txtValue = li[i].textContent || li[i].innerText;
        li[i].style.display = input.value.length && txtValue.toUpperCase().indexOf(filter) > -1 ? "block" : "none";
    }
}