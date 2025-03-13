function refreshData() {
    fetch('/api/stats')
        .then(response => response.json())
        .then(data => {
            document.getElementById("cpu-usage").textContent = data.SystemInfo.CPUUsage;
            document.getElementById("memory-usage").textContent = data.SystemInfo.MemoryUsage;
            document.getElementById("goroutines").textContent = data.SystemInfo.GoRoutines;
            document.getElementById("active-connections").textContent = data.Network.ActiveConnections;
            
            let ports = data.Network.OpenPorts.join(", ");
            document.getElementById("open-ports").textContent = ports;

            let processList = document.getElementById("process-list");
            processList.innerHTML = "";
            data.Processes.forEach(proc => {
                let row = `<tr>
                    <td>${proc.PID}</td>
                    <td>${proc.Name}</td>
                    <td>${proc.CPU}</td>
                </tr>`;
                processList.innerHTML += row;
            });
        })
        .catch(error => console.error("Error . Kondisi Fetch Data:", error));
}
    
setInterval(refreshData, 5000);
