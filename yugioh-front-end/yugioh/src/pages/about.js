import Image from "next/image";
import DMG from "../../public/DMG.webp";

const About = () => {
  return (
    <figure className="md:flex bg-slate-100 rounded-xl p-8 md:p-0 dark:bg-slate-800:">
      <div className="flex flex-wrap justify-center-left">
        <div className="w-6/12 sm:w-4/12 px-4">
          <Image
            src={DMG}
            alt="/"
            className="shadow rounded-full max-w-full h-auto align-middle border-none"
          />
          <div className="pt-6 md:p-8 text-center md:text-left space-y-4">
            <blockquote>
              <p className="text-lg font-medium">"INSERT ABOUT INFO HERE"</p>
            </blockquote>
          </div>
        </div>
      </div>
    </figure>
  );
};

export default About;
