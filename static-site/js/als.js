document.getElementById('al-dropdown').addEventListener("change", e => {

  const formData = new FormData(document.getElementById("als-form"));

  fetch("/api/setal", {
    method: "POST",
    body: formData,
  })
  .then(r => {

    if (r.ok) {
      dialang.session.al = formData.get("al");
      dialang.switchState("legend");
    } else {
      console.error("Failed to set Admin Language");
    }
  });
});
