
Set-Variable -Name CheemitPath -Value "$env:APPDATA\Cheemit\image"
Set-Variable -Name CheemitPathPrograms -Value "C:\Program Files\DiegoMagdaleno\cheemit"

if ( -not (Test-Path -Path $CheemitPath -PathType Container)) {
    try {
        New-Item -Path $CheemitPath -ItemType Directory -ErrorAction Stop | Out-Null
    } catch {
        Write-Error -Message "Unable to create direcotry '$CheemitPath'. Error was $_" -ErrorAction Stop
    }
    "Successfully created the directory for cheemit! '$CheemitPath."
} else {
    "Directory for cheemit already existed at $CheemitPath, will update contents of the directory"
}

Copy-Item -Path ".\resources\images\*" -Destination $CheemitPath -Recurse
"Preparting to install cheemit... Building necessary directories."

if ( -not (Test-Path -Path $CheemitPathPrograms -PathType Container )) {
    try {
        New-Item -Path $CheemitPathPrograms -ItemType Directory -ErrorAction Stop | Out-Null
    } catch {
        Write-Error -Message "Unable to create directory '$CheemitPathPrograms'. Error was $_" -ErrorAction Stop
    }
    "Successfully created the directory in programs for Cheemit! '$CheemitPathPrograms'"
} else {
    "Directory for cheemit in programs already existed at $CheemitPathPrograms, will update contents of the directory after build."
}

go build .

Copy-Item -Path ".\cheemit.exe" -Destination "$CheemitPathPrograms\cheemit.exe"