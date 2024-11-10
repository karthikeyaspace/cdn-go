document
  .getElementById("uploadForm")
  .addEventListener("submit", async function (e) {
    e.preventDefault();

    const indexFile = document.getElementById("indexFile").files[0];
    const styleFile = document.getElementById("styleFile").files[0];
    const formData = new FormData();

    formData.append("index.html", indexFile);
    formData.append("style.css", styleFile);

    try {
      const response = await fetch("http://localhost:8080/upload", {
        method: "POST",
        body: formData,
      });

      if (!response.ok) {
        throw new Error("Failed to upload files");
      }

      const result = await response.json();
      document.getElementById(
        "response"
      ).innerHTML = `Website deployed! visit: <a href="http://localhost:8080/view/${result.key}" target="_blank">http://localhost:8080/view/${result.key}</a>`;
    } catch (error) {
      document.getElementById("response").innerHTML = `Error: ${error.message}`;
    }
  });
