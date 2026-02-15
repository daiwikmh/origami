# UI Color Scheme Update Guide

## New Color Palette

```css
/* Primary Colors */
--black: #000000;
--dark-bg: #070709;
--purple: #E1C4E9;
--gray: #232323;
```

## Changes Needed

### 1. Remove Scanning Line Animation

**In `handlers/testing_html.go` and `handlers/dashboard_html.go`:**

**Remove this CSS:**
```css
.scan-line {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 2px;
    background: linear-gradient(90deg, transparent, #00ff9f, transparent);
    animation: scan 4s linear infinite;
    pointer-events: none;
    z-index: 9999;
}

@keyframes scan {
    0% { transform: translateY(0); }
    100% { transform: translateY(100vh); }
}
```

**Remove this HTML:**
```html
<div class="scan-line"></div>
```

### 2. Update Color Variables

**Replace these colors throughout:**

| Old Color | New Color | Usage |
|-----------|-----------|-------|
| `#00ff9f` (green) | `#E1C4E9` (purple) | Primary accent, text |
| `#ff006e` (pink) | `#E1C4E9` (purple) | Headers, buttons |
| `#8338ec` (purple) | `#232323` (gray) | Borders, secondary |
| `#3a86ff` (blue) | `#232323` (gray) | Input borders |
| `#0f172a` (navy) | `#070709` (dark) | Card backgrounds |
| `#1e293b` (dark blue) | `#232323` (gray) | Secondary backgrounds |

### 3. Specific Updates

**Body Background:**
```css
body {
    font-family: 'Orbitron', monospace;
    background: #000000;
    color: #E1C4E9;
    /* Remove the complex gradient background */
}
```

**Header:**
```css
header {
    background: #070709;
    border: 2px solid #E1C4E9;
    /* Remove glow effects */
}

h1 {
    color: #E1C4E9;
    /* Remove text-shadow */
}
```

**Cards:**
```css
.card {
    background: #070709;
    border: 1px solid #232323;
    /* Remove complex box-shadow */
}
```

**Buttons:**
```css
.btn {
    background: #E1C4E9;
    color: #000000;
    border: 2px solid #E1C4E9;
    /* Simplify hover effects */
}

.btn:hover {
    background: #232323;
    color: #E1C4E9;
}
```

**Inputs:**
```css
input, textarea {
    background: #000000;
    border: 2px solid #232323;
    color: #E1C4E9;
}

input:focus {
    border-color: #E1C4E9;
    /* Remove glow effects */
}
```

**Method Badges:**
```css
.method.get {
    background: #E1C4E9;
    color: #000000;
}

.method.post {
    background: #232323;
    color: #E1C4E9;
}
```

## Quick Find & Replace

Use these sed commands to quickly update colors:

```bash
# Navigate to handlers directory
cd /home/daiwi/inject/handlers

# Backup files first
cp testing_html.go testing_html.go.backup
cp dashboard_html.go dashboard_html.go.backup

# Replace green with purple
sed -i 's/#00ff9f/#E1C4E9/g' testing_html.go dashboard_html.go

# Replace pink with purple
sed -i 's/#ff006e/#E1C4E9/g' testing_html.go dashboard_html.go

# Replace specific purples with gray
sed -i 's/#8338ec/#232323/g' testing_html.go dashboard_html.go

# Replace blues with gray
sed -i 's/#3a86ff/#232323/g' testing_html.go dashboard_html.go

# Replace navy with dark
sed -i 's/#0f172a/#070709/g' testing_html.go dashboard_html.go

# Replace dark blue with gray
sed -i 's/#1e293b/#232323/g' testing_html.go dashboard_html.go
```

## Manual Cleanup Required

After running the sed commands, manually:

1. **Remove scan-line CSS** (search for `.scan-line`)
2. **Remove scan-line HTML** (search for `<div class="scan-line">`)
3. **Remove complex gradients** in headers
4. **Simplify box-shadow effects**
5. **Remove text-shadow effects**
6. **Test the UI** at http://localhost:8080/dashboard and /test

## Verification

After making changes:

```bash
# Rebuild
go build

# Test
./origami

# Visit:
# - http://localhost:8080/dashboard
# - http://localhost:8080/test
# - http://localhost:8080/docs
```

## Result

The new UI will have:
- ✅ Clean, minimal design
- ✅ Purple (#E1C4E9) as primary color
- ✅ Black/dark gray backgrounds
- ✅ No scanning animation
- ✅ Reduced visual noise
- ✅ Better readability
