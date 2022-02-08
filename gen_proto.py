from importlib.resources import path
from multiprocessing.spawn import _main
import os
from pathlib import Path
import chardet
import subprocess
import re
from typing import List, Optional

BASE_DIR_NAME = "source"
BASE_DIR = Path(BASE_DIR_NAME)

API_DIR_NAME = "tensorflow_serving/apis"
API_DIR = BASE_DIR.joinpath(API_DIR_NAME)


TARGET_BASE_DIR_NAME = "pbschema"
TARGET_BASE_DIR = Path(TARGET_BASE_DIR_NAME)

default_environ = dict(os.environ)
env = {}
env.update(default_environ)


def move_file(p: Path) -> None:
    """将pb文件处理后移动至指定文件夹

    Args:
        p (Path): 源文件位置
    """
    # package
    p_str = str(p)
    package_str = "/".join(p_str.split("/")[1:])
    go_package_str = "/".join(p_str.split("/")[1:-1])
    # 删除原有的go_package信息
    with open(p) as f:
        lines = f.readlines()
    without_go_package_lines = []
    for line in lines:
        if "option go_package" in line:
            print(f"****************** 文件{package_str}中有go_package,删除")
            continue
        else:
            if "go_package" in line:
                print(f"#############")
                print(line)
            without_go_package_lines.append(line)
    # 添加新的go_package信息
    wlines = []
    for line in without_go_package_lines:
        wlines.append(line)
        if line.startswith("package "):
            goopt = f'option go_package = "{go_package_str}";'
            wlines.append(goopt)

    # 写入
    target_file = TARGET_BASE_DIR.joinpath(package_str)
    if target_file.exists():
        print(f"xxxxxxxxxxxxx file {target_file} exists")
    else:
        print(f"------------- write to file {target_file}")
        if not target_file.parent.exists():
            print(f"------------- {target_file.parent} not exist,create")
            target_file.parent.mkdir(parents=True)
        with open(target_file, "w") as wf:
            wf.writelines(wlines)


def trans_package_to_path(package_line: str) -> Optional[Path]:
    """将pb的import行中指定的模块字符串转变为对应源文件路径.

    Args:
        package_line (str): import行字符串

    Returns:
        Optional[Path]: 源文件路径
    """
    r = re.search(r'"[\w|.|\-|/]+"', package_line)
    if r:
        s = r.group(0)
        p = BASE_DIR.joinpath(s.replace('"', '').strip())
        return p
    else:
        return None


def find_requirements(p: Path) -> List[Path]:
    """找到pb文件中需要的依赖项的路径.

    Args:
        p (Path): pb源文件

    Returns:
        List[Path]: 依赖项路径列表
    """
    packages: List[Path] = []
    with open(p) as f:
        lines = f.readlines()
    for line in lines:
        if 'import "tensorflow' in line:
            package_path = trans_package_to_path(line)
            if package_path:
                packages.append(package_path)
    return packages


def move_with_requirements(p: Path) -> None:
    """将指定pb源文件及其依赖移动至对应路径.

    Args:
        p (Path): 指定pb源文件路径
    """
    requirements: List[Path] = find_requirements(p)
    for req in requirements:
        move_with_requirements(req)
    move_file(p)


def genpb(p: Path, *pbp: Path) -> None:
    ppathstr = str(p)
    packagepath = "/".join(ppathstr.split("/")[1:])
    print(f"构造pb模块---{packagepath}")
    files = []
    for pb in pbp:
        pbpathstr = str(pb)
        filepath = "/".join(pbpathstr.split("/")[1:])
        files.append(filepath)

    files_str = " ".join(files)
    command = f"protoc -I pbschema --go_out=.  {files_str}"
    try:
        subprocess.run(command, capture_output=True, shell=True, check=True, env=env)
    except subprocess.CalledProcessError as ce:
        print(f"""命令:
        {command}
        执行失败""")
        if ce.stderr:
            encoding = chardet.detect(ce.stderr).get("encoding")
            content = ce.stderr.decode(encoding).strip()
        else:
            encoding = chardet.detect(ce.stdout).get("encoding")
            content = ce.stdout.decode(encoding).strip()
        print(content)
        raise ce
    except Exception as e:
        print(f"""命令:
        {command}
        执行失败""")
        print(f"error: {type(e)}")
        print(f"error_message: {str(e)}")
        raise e
    else:
        print(f"构造pb模块完成---{packagepath}")


def gengrpc(p: Path, *pbp: Path) -> None:
    ppathstr = str(p)
    packagepath = "/".join(ppathstr.split("/")[1:])
    print(f"构造grpc pb模块---{packagepath}")
    files = []
    for pb in pbp:
        pbpathstr = str(pb)
        filepath = "/".join(pbpathstr.split("/")[1:])
        files.append(filepath)

    files_str = " ".join(files)
    command = f"protoc -I pbschema --go_out=. --go-grpc_out=. {files_str}"
    print("start gen grpc pb")
    print(command)
    try:
        subprocess.run(command, capture_output=True, shell=True, check=True, env=env)
    except subprocess.CalledProcessError as ce:
        print(f"""命令:
        {command}
        执行失败""")
        if ce.stderr:
            encoding = chardet.detect(ce.stderr).get("encoding")
            content = ce.stderr.decode(encoding).strip()
        else:
            encoding = chardet.detect(ce.stdout).get("encoding")
            content = ce.stdout.decode(encoding).strip()
        print(content)
        raise ce
    except Exception as e:
        print(f"""命令:
        {command}
        执行失败""")
        print(f"error: {type(e)}")
        print(f"error_message: {str(e)}")
        raise e
    else:
        print(f"构造grpc pb模块完成---{packagepath}")


def itergen(p: Path) -> None:
    print("开始创建pb的go模块")
    if p.is_dir():
        pbs = []
        dirs = []
        for cp in p.iterdir():
            if cp.is_dir():
                dirs.append(cp)
            else:
                if cp.suffix == ".proto":
                    pbs.append(cp)
        print(f"文件夹{p}下有{len(pbs)}个pb文件和{len(dirs)}个文件夹")
        if len(dirs):
            for i in dirs:
                print(f"进入文件夹{i}")
                itergen(i)
        if len(pbs) > 0:
            if p.name == "apis":
                print("构造grpc模块")
                gengrpc(p, *pbs)
                pass
            else:
                genpb(p, *pbs)


def main() -> None:
    # 移动源文件到pbschema
    for f in API_DIR.iterdir():
        if f.is_file() and f.suffix == ".proto":
            move_with_requirements(f)

    # 生成go文件
    itergen(TARGET_BASE_DIR)


if __name__ == "__main__":
    main()
