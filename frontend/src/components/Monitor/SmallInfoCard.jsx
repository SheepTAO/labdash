const SmallInfoCard = ({ label, value, icon, color }) => {
    const colors = {
        indigo: "bg-indigo-50 text-indigo-600 border-indigo-100",
        amber: "bg-amber-50 text-amber-600 border-amber-100",
        emerald: "bg-emerald-50 text-emerald-600 border-emerald-100",
        rose: "bg-rose-50 text-rose-600 border-rose-100",
    };
    return (
        <div className={`rounded-xl px-3 py-2 border flex items-center justify-between ${colors[color]}`}>
            <div className="flex items-center gap-2 opacity-80">
                {icon}
                <span className="text-[10px] font-bold uppercase tracking-wider">{label}</span>
            </div>
            <div className="text-base font-black tracking-tight text-slate-800">{value}</div>
        </div>
    )
}

export default SmallInfoCard;
