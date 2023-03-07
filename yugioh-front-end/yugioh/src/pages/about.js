import Image from "next/image";
import DMG from "../../public/DMG.webp";

const About = () => {
  return (
    <div className="flex flex-row bg-slate-100 rounded-xl p-8 dark:bg-slate-800:">
        <div className="flex basis-1/4">
          <Image
            src={DMG}
            alt="/"
            className="max-h-full w-3/4 px-4 shadow rounded-full h-auto align-middle border-none"
          />
            </div>
          <div className="basis-1/4">
            <blockquote>
              <p className="text-lg font-medium">"INSERT ABOUT INFO HERE"</p>
            </blockquote>
          </div>
        </div>
  );
};

export default About;
