import { motion } from 'framer-motion';

// --- Component: Floating Orb ---
const FloatingOrb = ({ positionClass, animatePath, durations, colorClass, glowColorRGB, isHovered }) => (
  <motion.div 
    className={`absolute top-1/2 -mt-[200px] -ml-[200px] ${positionClass}`}
    animate={animatePath}
    transition={{ 
      x: { duration: durations.x, repeat: Infinity, ease: "easeInOut" },
      y: { duration: durations.y, repeat: Infinity, ease: "easeInOut" }
    }}
  >
     <motion.div 
       variants={{ 
         hover: { 
           scale: 0.25, 
           opacity: 1, 
           filter: "blur(0px)",
           borderWidth: "6px",
           borderColor: "rgba(255,255,255,0.9)",
           boxShadow: `0 0 50px 15px rgba(${glowColorRGB}, 0.6)` 
         },
         idle: { 
           scale: 1, 
           opacity: 0.3, 
           filter: "blur(80px)",
           borderWidth: "0px",
           borderColor: "rgba(255,255,255,0)",
           boxShadow: `0 0 0 0 rgba(${glowColorRGB}, 0)`
         }
       }}
       initial="idle"
       animate={isHovered ? "hover" : "idle"}
       transition={{ duration: 0.4, ease: "easeOut" }}
       className={`w-[400px] h-[400px] rounded-full mix-blend-multiply ${colorClass}`}
     />
  </motion.div>
);

export default FloatingOrb;
