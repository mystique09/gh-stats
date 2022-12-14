export function buildResult({ label, color }: { label: string, color: string }) {
  return `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="110" height="20">
  <linearGradient id="smooth" x2="0" y2="100%">
    <stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
    <stop offset="1" stop-opacity=".1"/>
  </linearGradient>

  <mask id="round">
    <rect width="110" height="20" rx="3" fill="#fff"/>
  </mask>

  <g mask="url(#round)">
    <rect width="83" height="20" fill="#555"/>
    <rect x="83" width="27" height="20" fill="#${color}"/>

    <rect width="110" height="20" fill="url(#smooth)"/>
  </g>

  <g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11">
    <text x="42.5" y="15" fill="#010101" fill-opacity=".3">Profile Visits</text>
    <text x="42.5" y="14">${label}</text>
    <text x="95.5" y="15" fill="#010101" fill-opacity=".3">28</text>
    <text x="95.5" y="14">0</text>
  </g>
</svg>`;
}
