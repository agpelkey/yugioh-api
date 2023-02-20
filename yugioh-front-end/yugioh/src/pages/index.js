import Head from "next/head";
import Layout, { siteTitle } from "../../components/layout";
import utilStyles from "../../styles/utils.module.css";

export default function Home() {
  return (
    <Layout home>
      <Head>
        <title>{siteTitle}</title>
      </Head>
      <section className={utilStyles.headingMd}>
        <p>This is a simple web app to fetch various yugioh cards</p>
      </section>
    </Layout>
  );
}
